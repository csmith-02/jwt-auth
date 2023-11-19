package database

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables.")
	}
}
