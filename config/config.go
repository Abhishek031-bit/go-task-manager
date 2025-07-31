package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWT_SECRET string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatal("JWT_SECRET is not set in .env")
	}
}
