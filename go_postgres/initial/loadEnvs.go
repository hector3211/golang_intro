package initial

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvs() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
