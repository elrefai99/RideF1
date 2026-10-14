package emails

import (
	"time"

	"github.com/google/uuid"
)

type EmailsEnum string

const (
	EmailSent              EmailsEnum = "SENT"
	EmailDelivered         EmailsEnum = "DELIVERED"
	EmailDelivered_delayed EmailsEnum = "DELIVERED_DELAYED"
	EmailFailed            EmailsEnum = "FAILED"
)

type Emails struct {
	ID        uint       `gorm:"primaryKey"`
	Uuid      uuid.UUID  `gorm:"type:uuid;not null;index"`
	Status    EmailsEnum `gorm:"varchar(100);not null;index:idx_emails_status_email,priority:1;default:'SENT'"`
	Email_id  string     `gorm:"not null;index:idx_emails_status_email,priority:2"`
	Sender    string     `gorm:"not null;"`
	Receiver  string     `gorm:"not null;"`
	Subject   string     `gorm:"not null"`
	Body      string     `gorm:"not null"`
	IsClicked bool       `gorm:"not null;default:false"`
	IsOpened  bool       `gorm:"not null;default:false"`
	OpenedAt  time.Time
	ClickedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
