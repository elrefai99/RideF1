package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db, nil
}
