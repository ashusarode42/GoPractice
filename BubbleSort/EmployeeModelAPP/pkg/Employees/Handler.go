package Employees

import (
	database "EmployeeModelAPP/pkg/Database"
	"encoding/json"
	"net/http"
)

var employee Employee

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got Employee 1")

}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&employee)
	addEmployee(&employee)
	println("Added Employee to DB..")
	json.NewEncoder(w).Encode("Added Employee to cassandra db")

}

func addEmployee(employee *Employee) {
	query := `INSERT INTO employee(id,name,address,specification) values (?,?,?,?)`
	database.ExecuteQuery(query, employee.Id, employee.Name, employee.Address, employee.Specification)
}
