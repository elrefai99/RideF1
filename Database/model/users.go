package models

import (
	"time"

	"github.com/google/uuid"
)

// ENUMS
type (
	UserEnum   string
	DriverEnum string
	EmailsEnum string
	OTPEnum    string
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

const (
	EmailSent              EmailsEnum = "SENT"
	EmailDelivered         EmailsEnum = "DELIVERED"
	EmailDelivered_delayed EmailsEnum = "DELIVERED_DELAYED"
	EmailFailed            EmailsEnum = "FAILED"
)

const (
	OTPSuccess OTPEnum = "SUCCESS"
	OTPReject  OTPEnum = "REJECT"
	OTPPending OTPEnum = "PENDING"
	OTPFailed  OTPEnum = "FAILED"
)

// User SCHEMA
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Uuid      uuid.UUID `gorm:"varchar(255);not null;index:idx_users_status_uuid,priority:1"`
	Fullname  string    `gorm:"varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null;index:idx_users_status_email,priority:1"`
	Password  string    `gorm:"varchar(255);not null"`
	Status    UserEnum  `gorm:"type:user_status;not null;default:'INACTIVE';index:idx_users_status_email,priority:2;index:idx_users_status_username,priority:1;index:idx_users_status_uuid,priority:2"`
	Username  string    `gorm:"type:varchar(255);not null;index:idx_users_status_username,priority:2"`
	Avatar    string    `gorm:"varchar(255);not null"`
	Google_id string    `gorm:"varchar(255);not null"`
	Phone     string    `gorm:"varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Rider_Profiles  []Rider_Profile  `gorm:"foreignKey:Rider_id"`
	Driver_Profiles []Driver_Profile `gorm:"foreignKey:Driver_id"`
	Token           []Tokens         `gorm:"foreignKey:User_id"`
	EmailUser       []Emails         `gorm:"foreignKey:User_id"`
	OTPUser         []OTP            `gorm:"foreignKey:User_id"`
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
	Uuid           uuid.UUID  `gorm:"type:uuid;not null;index"`
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
	ID           uint      `gorm:"primaryKey"`
	Uuid         uuid.UUID `gorm:"type:uuid;not null;index"`
	User_id      uint      `gorm:"null;index"`
	RefrashToken string    `gorm:"varchar(255);not null"`
	IPAddress    string    `gorm:"type:varchar(255);not null;default:''"`
	FCM          string    `gorm:"varchar(255);not null"`
	ExpireAt     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time

	TokensUser User `gorm:"foreignKey:User_id;references:ID"`
}

type OTP struct {
	ID        uint      `gorm:"primaryKey"`
	Uuid      uuid.UUID `gorm:"type:uuid;not null;index"`
	User_id   uint      `gorm:"null;index:idx_opts_status_userid,priority:1"`
	Status    OTPEnum   `gorm:"varchar(50);not null;index:idx_opts_status_userid,priority:2;default:'PENDING'"`
	Sender_by string    `gorm:"varchar(50);not null"`
	Message   string    `gorm:"varchar(255);not null"`
	ExpireAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time

	OTP_user User `gorm:"foreignKey:User_id;references:ID"`
}

type Emails struct {
	ID        uint       `gorm:"primaryKey"`
	Uuid      uuid.UUID  `gorm:"type:uuid;not null;index"`
	User_id   uint       `gorm:"null"`
	Status    EmailsEnum `gorm:"varchar(100);not null;index:idx_emails_status_email,priority:1"`
	Email_id  string     `gorm:"not null;index:idx_emails_status_email,priority:2"`
	Sender    string     `gorm:"not null;"`
	Receiver  string     `gorm:"not null;"`
	Subject   string     `gorm:"not null"`
	Body      string     `gorm:"not null"`
	IsClicked bool       `gorm:"not null;"`
	IsOpened  bool       `gorm:"not null;"`
	OpenedAt  time.Time
	ClickedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time

	EmailUser User `gorm:"foreignKey:User_id;references:ID"`
}
