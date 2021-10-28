package controllers

import (
	"encoding/json"
	"go_crud/database"
	"go_crud/entity"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllEmployee get all Employee data
func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	var employees []entity.Employee
	database.Connector.Find(&employees)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

//GetEmployeeByID returns Employee with specific ID
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var employee entity.Employee
	database.Connector.First(&employee, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

//CreateEmployee creates Employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var employee entity.Employee
	json.Unmarshal(requestBody, &employee)

	database.Connector.Create(employee)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}

//UpdateEmployeeByID updates Employee with specific ID
func UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var employee entity.Employee
	json.Unmarshal(requestBody, &employee)
	database.Connector.Save(&employee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

//DeletEmployeeByID delete Employee with specific ID
func DeletEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var employee entity.Employee
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&employee)
	w.WriteHeader(http.StatusNoContent)
}
