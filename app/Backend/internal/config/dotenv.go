package config

import (
	"os"

	"github.com/joho/godotenv"
)

type IDotenv struct {
	PORT             string
	DATABASE_URI     string
	DATABASE         string
	ACCESS_TOKEN_JWT string
}

func DotenvConfig() (*IDotenv, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	envData := &IDotenv{
		PORT:             os.Getenv("PORT"),
		DATABASE_URI:     os.Getenv("DATABASE_URI"),
		DATABASE:         os.Getenv("DATABASE"),
		ACCESS_TOKEN_JWT: os.Getenv("ACCESS_TOKEN_JWT"),
	}

	return envData, nil
}
