package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadDotENV adalah fungsi untuk meload .env
func LoadDotENV() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
