package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading environment file...")
	}

	return os.Getenv(key)
}