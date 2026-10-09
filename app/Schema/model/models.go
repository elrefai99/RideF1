package main

import "time"

// User SCHEMA
type UserEnum string

const (
	UserActive   UserEnum = "ACTIVE"
	UserINActive UserEnum = "INACTIVE"
	UserBLOCK    UserEnum = "BLOCK"
	UserBAN      UserEnum = "BAN"
)

type User struct {
	ID        int      `gorm:"primaryKey"`
	Uuid      string   `grom:"varchar(255);not null;index:idx_users_status_uuid,priority:2"`
	Fullname  string   `grom:"varchar(255);not null"`
	Email     string   `grom:"varchar(255);not nullindex:idx_users_status_email,priority:2"`
	Password  string   `grom:"varchar(255);not null"`
	Status    UserEnum `grom:"type:user_status;not null;varchar(255);default:'INACTIVE';index:idx_users_status_email,priority:1;index:idx_users_status_username,priority:1"`
	Username  string   `grom:"varchar(255);not null;index:idx_users_status_username,priority:2"`
	Avatar    string   `grom:"varchar(255);not null"`
	Google_id string   `grom:"varchar(255);not null"`
	Phone     string   `grom:"varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
