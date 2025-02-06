package service

import (
	"employeeeDirectory/models"
	"net/http"
)

type EmployeeService interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request)                       //Create
	GetEmployee(w http.ResponseWriter, r *http.Request) (models.Employee, error) //Read
	UpdateEmployee(w http.ResponseWriter, r *http.Request) error                 //Update
	DeleteEmployee(w http.ResponseWriter, r *http.Request) error                 //Delete
	ListAllEmployees()
}
