package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if os.Getenv("APP_ENV") == "live" {
		return nil // Production gets configuration from the process environment.
	}

	envFile := ".env"
	if os.Getenv("APP_ENV") == "development" {
		envFile = ".env.dev"
	}

	return godotenv.Load(envFile)
}
