package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SecretKey string
)

func LoadConfig() {
	// Attempt to load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables directly")
	}

	// Load environment variables
	SecretKey = os.Getenv("JWT_SECRET")

	if SecretKey == "" {
		log.Fatal("JWT_SECRET not set in environment variables")
	}
}
