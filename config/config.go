package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWT_SECRET string
var DATABASE_URL string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatal("JWT_SECRET is not set in .env")
	}
	DATABASE_URL = os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}
}
