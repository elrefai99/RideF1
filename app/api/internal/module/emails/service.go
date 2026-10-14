package emails

import (
	"context"
	"errors"

	"github.com/elrefai99/RideF1/app/api/internal/Common/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IEmailSerice struct {
	db     *gorm.DB
	schema *repository.IRepository[Emails]
}

type IPayload struct {
	User_id  uint
	Email_id string
	Sender   string
	Receiver string
	Subject  string
	Body     string
}

func EmailSerice(db *gorm.DB) *IEmailSerice {
	schema := repository.Repository[Emails](db)

	return &IEmailSerice{
		db:     db,
		schema: schema,
	}
}

func (e *IEmailSerice) FindEmail(ctx context.Context, payload IPayload) (any, error) {

	query := map[string]any{
		"Status":   EmailSent,
		"Email_id": payload.Email_id,
	}

	data, err := e.schema.FindOne(ctx, query)
	if err == nil {
		return false, errors.New("Username or Email is already used")
	}

	return data, nil
}

func (e *IEmailSerice) CreateNewEmail(ctx context.Context, payload IPayload) (bool, error) {
	createEmail := &Emails{
		Uuid:     uuid.New(),
		Status:   EmailSent,
		Email_id: payload.Email_id,
		Sender:   payload.Sender,
		Receiver: payload.Receiver,
		Subject:  payload.Subject,
		Body:     payload.Body,
	}

	email := []Emails{*createEmail}

	if err := e.schema.Create(ctx, &email); err != nil {
		return false, err
	}

	return true, nil
}

func (e *IEmailSerice) OpenEmail() {

}

func (e *IEmailSerice) ClickEmail() {

}
