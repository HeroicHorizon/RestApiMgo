package employee

import (
	"time"
)

// Employee employeestruct
// swagger:model
type Employee struct {
	EmpID       string    `json:"empid" bson:"empid"`
	FName       string    `json:"fname"`
	LName       string    `json:"lname"`
	Designation string    `json:"designation"`
	Phone       int64     `json:"phone"`
	JoiningDate time.Time `json:"joiningdate"`
}
