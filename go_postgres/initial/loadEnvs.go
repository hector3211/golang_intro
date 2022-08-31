package initial

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvs() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
