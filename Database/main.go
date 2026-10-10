package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elrefai99/RideF1/Database/internal/migrations"
	models "github.com/elrefai99/RideF1/Database/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	_ = godotenv.Load()

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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

	err = db.AutoMigrate(
		&models.User{},
		&models.Rider_Profile{},
		&models.Driver_Profile{},
		&models.Tokens{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done Push Database nigga")
}
