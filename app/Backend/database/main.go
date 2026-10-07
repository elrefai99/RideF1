package main

import (
	"os"

	"github.com/joho/godotenv"
)

type IDotenv struct {
	PORT             string
	DATABASE_URL     string
	DATABASE         string
	ACCESS_TOKEN_JWT string
}

func DotenvConfig() (*IDotenv, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	envData := &IDotenv{
		DATABASE_URL:     os.Getenv("DATABASE_URL"),
	}

	return envData, nil
}

func main() {

}
