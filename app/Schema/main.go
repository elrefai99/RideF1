package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elrefai99/PBE/app/Schema/migrations"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	if err := migrations.EnumCheck(db); err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done Push Database nigga")
}
