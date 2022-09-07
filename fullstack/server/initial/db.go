package initial

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// declare global variable
var DB *gorm.DB

func ConnectDB() {
	var err error
	// load our env for our db
	dsn := os.Getenv("DB_URL")
	// establish connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if error retrun fatal
	if err != nil {
		log.Fatal("Failed to connect database")
	}
}
