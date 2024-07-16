package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN       string
	JWTSecret   string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Detect if running in Docker
	inDocker := RunningInDocker()

	// Use different DSN based on the environment
	var dsn string
	if inDocker {
		dsn = os.Getenv("DOCKER_DB_DSN")
	} else {
		dsn = os.Getenv("LOCAL_DB_DSN")
	}

	return &Config{
		DBDSN:     dsn,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

// RunningInDocker function checks if the app is running in Docker
func RunningInDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}
