package main

import (
	"go_postgres/initial"
	"go_postgres/models"
)

func init() {
	initial.LoadEnvs()
	initial.ConnectDB()
}

func main() {
	// Migrate the schema
	initial.DB.AutoMigrate(&models.Product{})
}
