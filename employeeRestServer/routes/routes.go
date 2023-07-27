package routes

import (
	emp "employeeRestServer/employeeServer"
	"mux"
)

type Router struct {
}

func (r *Router) Routes() *mux.Router {

	myRouter := mux.NewRouter().StrictSlash(true)
	e := emp.EmployeeServer{}
	// this func is to make db connection avaialble to all routes within the routHandler methods
	
	myRouter.HandleFunc("/employees", e.AddEmployees).Methods("POST")
	myRouter.HandleFunc("/employees", e.ListEmployees).Methods("GET")

	return myRouter
}
