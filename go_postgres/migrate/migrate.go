package main

import (
	"go_postgres/initial"
	"go_postgres/models"
)

// init fucntion inside a package main runs first
func init() {
	// load our environment variables
	initial.LoadEnvs()
	// establish a connection to our db
	initial.ConnectDB()
}

func main() {
	// Migrate our Product model to db
	initial.DB.AutoMigrate(&models.Product{})
}
