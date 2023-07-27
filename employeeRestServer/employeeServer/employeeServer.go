package employeeServer

import (
	dbCon "employeeRestServer/databaseConnection"
	model "employee_rest_server/Models"
	"log"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//employee server actions
type EmployeeServer struct{}

func (e *EmployeeServer) AddEmployees(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee model.Employee
	json.Unmarshal(reqBody, &employee)

	var dbC dbCon.DatabaseConnection
	db, err := dbC.GetConnection()

	if err != nil {
		fmt.Println("Could not connect to Database", err)

		return
	}

	employee.Address = employee.Address
	json.NewEncoder(w).Encode(&employee)
	db.Create(&employee)
	fmt.Println(w)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}

func (e *EmployeeServer) ListEmployees(w http.ResponseWriter, r *http.Request) {
	var dbC dbCon.DatabaseConnection
	db, err := dbC.GetConnection()

	if err != nil {
		fmt.Println("Could not connect to Database", err)

		return
	}
	var employee model.Employee
	var employees []model.Employee
	fmt.Println(employees)
	result := db.Preload("Address").Find(&employee)

	rows, error := result.Rows()
	if error != nil {
		log.Fatal("could not fetch all employees " + error.Error())
		return
	}
	for rows.Next() {
		rows.Scan(employee)

		fmt.Println(employee)
		employees = append(employees, employee)

	}
	fmt.Println(employees)
	json.NewEncoder(w).Encode(&employees)
}
