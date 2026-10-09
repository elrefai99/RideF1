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
	Uuid      string   `grom:"not null;index:idx_users_status_uuid,priority:2"`
	Fullname  string   `grom:"not null"`
	Email     string   `grom:"not nullindex:idx_users_status_email,priority:2"`
	Password  string   `grom:"not null"`
	Status    UserEnum `grom:"type:user_status;not null;varchar(255);default:'INACTIVE';index:idx_users_status_email,priority:1;index:idx_users_status_username,priority:1"`
	Username  string   `grom:"not null;index:idx_users_status_username,priority:2"`
	Avatar    string   `grom:"not null"`
	Google_id string   `grom:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CoHost struct {
}

// Properties Schema

type PropertiesEnum string

const (
	PropertyActive  PropertiesEnum = "ACTIVE"
	PropertyPending PropertiesEnum = "PENDING"
	PropertyDelete  PropertiesEnum = "DELETED"
	PropertyStop    PropertiesEnum = "STOP"
)

type Properties struct {
}

type PropertiesImages struct {
}
