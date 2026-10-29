package main

import (
	"fmt"
	"net/http"

	"github.com/RestApiMgo/propparser"

	"github.com/RestApiMgo/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//go:generate swagger generate spec -m -o ./swagger.json
func main() {
	config := propparser.ReadConfig()
	var port = ":" + config.Port
	controller.Init()
	//init router
	router := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "QUERY", "OPTIONS"})
	router.HandleFunc("/v1/Employees", controller.GetAllEmployees).Methods("GET", "OPTION")
	router.HandleFunc("/v1/Employee/{fname}", controller.GetEmployee).Methods("GET", "OPTION")
	router.HandleFunc("/v1/Employee", controller.InsertEmployee).Methods("POST", "OPTION")
	router.HandleFunc("/v1//Employee/{fname}", controller.UpdateEmployee).Methods("PUT", "OPTION")
	router.HandleFunc("/v1/Employee/{fname}", controller.DeleteEmployee).Methods("DELETE", "OPTION")
	router.HandleFunc("/v1/search/{empid}", controller.SearchEmployee).Methods("QUERY")

	fmt.Printf("application listening port%s\n", port)
	http.ListenAndServe(port, handlers.CORS(corsObj, corsMethods)(router))

}
