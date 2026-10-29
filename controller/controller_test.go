package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/RestApiMgo/employee"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1/search/{empid}", SearchEmployee).Methods("QUERY")
	return router
}

func stubLookup(t *testing.T, emp employee.Employee) {
	old := findByEmpID
	t.Cleanup(func() { findByEmpID = old })
	findByEmpID = func(empid string) (employee.Employee, error) {
		if empid == emp.EmpID {
			return emp, nil
		}
		return employee.Employee{}, mgo.ErrNotFound
	}
}

func TestSearchEmployeeFound(t *testing.T) {
	want := employee.Employee{
		EmpID: "E101", FName: "Ann", LName: "Lee", Designation: "Dev",
		Phone: 123456, JoiningDate: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
	}
	stubLookup(t, want)

	rec := httptest.NewRecorder()
	newRouter().ServeHTTP(rec, httptest.NewRequest("QUERY", "/v1/search/E101", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var got employee.Employee
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestSearchEmployeeNotFound(t *testing.T) {
	stubLookup(t, employee.Employee{EmpID: "E101"})

	rec := httptest.NewRecorder()
	newRouter().ServeHTTP(rec, httptest.NewRequest("QUERY", "/v1/search/E999", nil))

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}

func TestSearchEmployeeGetNotRouted(t *testing.T) {
	stubLookup(t, employee.Employee{EmpID: "E101"})

	rec := httptest.NewRecorder()
	newRouter().ServeHTTP(rec, httptest.NewRequest("GET", "/v1/search/E101", nil))

	if rec.Code == http.StatusOK {
		t.Fatalf("GET must not reach the handler, got %d", rec.Code)
	}
}
