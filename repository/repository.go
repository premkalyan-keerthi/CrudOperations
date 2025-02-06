package repository

import (
	"context"
	"employeeeDirectory/db"
	"employeeeDirectory/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var employeeCollection *mongo.Collection

var id int = 1000

func InitEmplloyeeRepository(coll *mongo.Collection) {
	employeeCollection = coll
}

func CreateEmployee(ctx context.Context, employee models.Employee) error {
	_, err := employeeCollection.InsertOne(ctx, employee)

	return err
}

type EmployeeRepo struct {
	employees map[int]models.Employee
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{employees: make(map[int]models.Employee)}
}

func (r *EmployeeRepo) CreateEmployee(w http.ResponseWriter, req *http.Request) {

	defer func() {

		if e := recover(); e != nil {
			fmt.Println("I got a panick!!", e)
		}
	}()

	w.Header().Set("Content-Type", "application/json")

	var employee models.Employee

	err := json.NewDecoder(req.Body).Decode(&employee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	employee.EmployeeID = id

	id++

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.GetCollection("employeedirectory", "employees")
	_, err = collection.InsertOne(ctx, employee)

	if err != nil {
		http.Error(w, "Failed to save to DB", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

}

func (r *EmployeeRepo) GetEmployee(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	id := params["id"]

	var employee models.Employee

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.GetCollection("employeedirectory", "employees")
	intval, _ := strconv.Atoi(id)
	err := collection.FindOne(ctx, bson.M{"employeeID": intval}).Decode(&employee)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to save to DB", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employee)

}

func (r *EmployeeRepo) UpdateEmployee(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	id := params["id"]

	var employee models.Employee

	err := json.NewDecoder(req.Body).Decode(&employee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": employee}

	collection := db.GetCollection("employeedirectory", "employees")

	intval, _ := strconv.Atoi(id)

	_, err = collection.UpdateOne(ctx, bson.M{"employeeID": intval}, update)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to save to DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (r *EmployeeRepo) DeleteEmployee(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	id := params["id"]

	var employee models.Employee

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.GetCollection("employeedirectory", "employees")
	intval, _ := strconv.Atoi(id)
	_, err := collection.DeleteOne(ctx, bson.M{"employeeID": intval})

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to save to DB", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employee)
}

func (r *EmployeeRepo) ListAllEmployees() {
	fmt.Println(r.employees)
}
