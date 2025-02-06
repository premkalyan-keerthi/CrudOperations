package main

import (
	"employeeeDirectory/db"
	"employeeeDirectory/repository"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	db.Connect()

	router := mux.NewRouter()

	rep := repository.NewEmployeeRepo()

	router.HandleFunc("/employees", rep.CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}", rep.GetEmployee).Methods(http.MethodGet)

	router.HandleFunc("/employees/{id}", rep.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/employees/{id}", rep.DeleteEmployee).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)
	fmt.Println("Starting Server")
}
