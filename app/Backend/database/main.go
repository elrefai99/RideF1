package main

import (
	"log"
	"os"

	"github.com/elrefai99/PBE/app/database/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDotenv struct {
	PORT             string
	DATABASE_URL     string
	DATABASE         string
	ACCESS_TOKEN_JWT string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Emails{},
	)

	if err != nil {
		log.Fatal(err)
	}

}
