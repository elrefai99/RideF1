package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elrefai99/RideF1/app/api/internal/database/migrations"
	"github.com/elrefai99/RideF1/app/api/internal/module/emails"
	"github.com/elrefai99/RideF1/app/api/internal/module/user"
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
		&user.User{},
		&user.Rider_Profile{},
		&user.Driver_Profile{},
		&user.Tokens{},
		&user.OTP{},
		&emails.Emails{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done Push Database nigga")
}
