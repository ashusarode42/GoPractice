package databaseConnection

import (
	"gorm"
	_ "gorm/dialects/postgres"
	"os"
	
)

// Database Connection struc
type DatabaseConnection struct {
	Db  *gorm.DB
	Err error
}



//get the existing DB Connection
func (con DatabaseConnection) GetConnection() (*gorm.DB, error) {
	if con.Db == nil {
		hostname:=os.Getenv("HOSTNAME")
		hostport:=os.Getenv("HOSTPORT")
		username:=os.Getenv("USRNAME")
		password:=os.Getenv("PASSWORD")
		con.Db, con.Err = con.Connect(hostname,hostport,username,password)
	}
	

	return con.Db, con.Err
 }

//connect to the database
func (con DatabaseConnection) Connect(hostname string, hostport string, username string, password string) (*gorm.DB, error) {

con.Db, con.Err = gorm.Open("postgres", " host="+hostname+" port="+hostport+" user="+username+" sslmode=disable password="+password)

	return con.Db, con.Err
}


