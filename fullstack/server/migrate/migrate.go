package main

import (
	"server/initial"
	"server/models"
)

func init() {
	initial.LoadEnvs()
	initial.ConnectDB()
}

func main() {
	initial.DB.AutoMigrate(&models.Book{})
}
