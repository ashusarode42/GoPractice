package main

import (
	dbCon "employeeRestServer/databaseConnection"
	route "employeeRestServer/routes"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	
	"net/http"
	
)

func main() {
	var db dbCon.DatabaseConnection
	// connect to database first
	err := godotenv.Load()
	if err != nil{
	   fmt.Println(err)
	}
	fmt.Println(os.Getenv("PASSWORD"))
	con, err := db.Connect(os.Getenv("HOSTNAME"),os.Getenv("HOSTPORT"),os.Getenv("USRNAME"),os.Getenv("PASSWORD"))
	defer con.Close()
	if err != nil {
		fmt.Println("Could not connect to Database", err)

		return
	}
	fmt.Println("database connection successfull")

	r := route.Router{}
	routes := r.Routes()
	error := http.ListenAndServe("localhost:1000", routes)
	if error != nil {
		fmt.Println("Exiting Server", error)

		return
	}

}
