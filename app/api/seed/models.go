package main

import "time"

type StatusEnum string
type IEmailType string

const (
	StatusEnumSent              StatusEnum = "SENT"
	StatusEnumDelivered         StatusEnum = "DELIVERED"
	StatusEnumDelivered_delayed StatusEnum = "DELIVERED_DELAYED"
	StatusEnumFailed            StatusEnum = "FAILED"
	StatusEnumNONE              StatusEnum = ""
)

const (
	IEmailTypeGmail  IEmailType = "gmail"
	IEmailTypeResend IEmailType = "resend"
)

type Emails struct {
	ID           int64      `gorm:"primaryKey"`
	Sender       string     `gorm:"not null;index:idx_emails_status_sender_receiver,priority:2"`
	Receiver     string     `gorm:"not null;index:idx_emails_status_sender_receiver,priority:3"`
	Subject      string     `gorm:"not null"`
	Body         string     `gorm:"not null"`
	Status       StatusEnum `gorm:"not null;index:idx_emails_status_sender_receiver,priority:1;index:idx_emails_status_is_clicked,priority:1;index:idx_emails_status_is_opened,priority:1;index:idx_emails_status_email_uuid,priority:1"`
	IsClicked    bool       `gorm:"not null;index:idx_emails_status_is_clicked,priority:2"`
	IsOpened     bool       `gorm:"not null;index:idx_emails_status_is_opened,priority:2"`
	Type         string     `gorm:"not null"`
	Email_id     string     `gorm:"not null;index:idx_emails_status_email_uuid,priority:2"`
	EmailType    IEmailType `gorm:"not null"`
	Uuid         string     `gorm:"not null;index:idx_emails_status_email_uuid,priority:3"`
	Broadcast_id string     `gorm:"not null"`
	Template_id  string     `gorm:"not null"`
	Bounce       struct {
		Message string
		SubType string
		Type    string
	} `gorm:"serializer:json;type:jsonb;not null"`
	Tags struct {
		Category string
	} `gorm:"serializer:json;type:jsonb;not null"`
	OpenedAt  time.Time
	ClickedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
