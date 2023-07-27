package Employees

import "github.com/go-chi/chi"

func EmployeeRoutes() *chi.Mux {

	router := chi.NewRouter()
	router.Get("/getEmployee", getHandler)
	router.Post("/addEmployee", postHandler)

	return router
}
