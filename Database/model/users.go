package models

import "time"

// ENUMS
type (
	UserEnum   string
	DriverEnum string
)

const (
	UserActive   UserEnum = "ACTIVE"
	UserINActive UserEnum = "INACTIVE"
	UserBLOCK    UserEnum = "BLOCK"
	UserBAN      UserEnum = "BAN"
)

const (
	DriverAccept  DriverEnum = "ACCEPT"
	DriverReject  DriverEnum = "REJECT"
	DriverPending DriverEnum = "PENDING"
	DriverBlock   DriverEnum = "BLOCK"
)

// User SCHEMA
type User struct {
	ID        uint     `gorm:"primaryKey"`
	Uuid      string   `grom:"varchar(255);not null;index:idx_users_status_uuid,priority:1"`
	Fullname  string   `grom:"varchar(255);not null"`
	Email     string   `gorm:"type:varchar(255);not null;index:idx_users_status_email,priority:1"`
	Password  string   `grom:"varchar(255);not null"`
	Status    UserEnum `gorm:"type:user_status;not null;default:'INACTIVE';index:idx_users_status_email,priority:2;index:idx_users_status_username,priority:1"`
	Username  string   `gorm:"type:varchar(255);not null;index:idx_users_status_username,priority:2"`
	Avatar    string   `grom:"varchar(255);not null"`
	Google_id string   `grom:"varchar(255);not null"`
	Phone     string   `grom:"varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Rider_Profiles  []Rider_Profile  `gorm:"foreignKey:Rider_id"`
	Driver_Profiles []Driver_Profile `gorm:"foreignKey:Driver_id"`
	Token           []Tokens         `gorm:"foreignKey:User_id"`
}

type Rider_Profile struct {
	ID        uint    `gorm:"primaryKey"`
	Uuid      string  `gorm:"varchar(255);not null"`
	Rider_id  uint    `gorm:"null;index"`
	Rating    float64 `gorm:"type:numeric(2,1);not null;default:0.0"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Rider User `gorm:"foreignKey:Rider_id;references:ID"`
}

type Driver_Profile struct {
	ID             uint       `gorm:"primaryKey"`
	Uuid           string     `gorm:"varchar(255);not null"`
	Driver_id      uint       `gorm:"index:idx_driver_profile,priority:1"`
	Status         DriverEnum `gorm:"type:varchar(255);not null;index:idx_driver_profile,priority:2;default:'PENDING'"`
	License_number string     `gorm:"varchar(255);not null"`
	License_Face   string     `gorm:"varchar(255);not null"`
	License_Back   string     `gorm:"varchar(255);not null"`
	Rating         float64    `gorm:"type:numeric(2,1);not null;default:0.0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	BlockAt        time.Time

	Driver User `gorm:"foreignKey:Driver_id;references:ID"`
}

type Tokens struct {
	ID           uint   `gorm:"primaryKey"`
	Uuid         string `gorm:"varchar(255);not null"`
	User_id      uint   `gorm:"null;index"`
	RefrashToken string `gorm:"varchar(255);not null"`
IPAddress string `gorm:"type:varchar(255);not null;default:''"`
	FCM          string `gorm:"varchar(255);not null"`
	ExpireAt     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time

	TokensUser User `gorm:"foreignKey:User_id;references:ID"`
}
