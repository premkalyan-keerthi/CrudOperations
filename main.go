package main

import (
	"employeeeDirectory/models"
	"employeeeDirectory/repository"
	"employeeeDirectory/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// CRUD Operations

func main() {

	repo := repository.NewEmployeeRepo()

	Execute(repo)

}

func Execute(repo service.EmployeeService) {

	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		//Create
		case http.MethodPost:
			{

				repo.CreateEmployee(w, r)

			}
		//Read
		case http.MethodGet:
			{
				path := r.URL.Query().Get("id")
				id, error := strconv.Atoi(path)
				fmt.Println(path)

				if error != nil {
					fmt.Println("Unable to convert to integer")
				}
				repo.GetEmployee(id)
			}
		//Update
		case http.MethodPut:
			{
				var emp models.Employee
				if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
					fmt.Println("Wrong data is sent")
				}
				repo.UpdateEmployee(emp)

			}
		//Delete
		case http.MethodDelete:
			{
				path := r.URL.Query().Get("id")
				fmt.Println(path)
				id, error := strconv.Atoi(path)

				if error != nil {
					fmt.Println("Unable to convert to integer")
				}
				repo.DeleteEmployee(id)
			}
		default:
			{
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}

		}
	})

	fmt.Println("Starting Server")

	http.ListenAndServe(":8080", nil)

}
