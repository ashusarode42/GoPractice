package Router

import (
	database "EmployeeModelAPP/pkg/Database"
	"EmployeeModelAPP/pkg/Employees"

	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func StartServer() *chi.Mux {
	database.SetupDBConnection()
	router := chi.NewRouter()
	router.Mount("/api/employees", Employees.EmployeeRoutes())
	fmt.Println("Server is listing on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
	return router
}
