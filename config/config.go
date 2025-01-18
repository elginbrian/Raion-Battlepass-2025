package config

import (
	"log"
	"os"
)

func GetServerPort() string {
	port := ":8080"
	return port
}

func GetDatabaseURL() string {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalf("DATABASE_URL environment variable is not set")
	}
	return dbURL
}

func GetJWTSecret() string {
    return "my-very-secure-secret"
}