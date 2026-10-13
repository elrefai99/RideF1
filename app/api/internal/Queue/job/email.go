package JobQueue

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/elrefai99/RideF1/app/api/internal/Queue/model"
	"github.com/elrefai99/RideF1/app/api/internal/module/emails"
	"github.com/resend/resend-go/v2"
	"gorm.io/gorm"
)

func SendEmail(jobs model.IEmailJob, db *gorm.DB) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", username, password, host)

	message := []byte(
		"From: " + username + "\r\n" +
			"To: " + jobs.To + "\r\n" +
			"Subject: " + jobs.Subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain;charset=UTF-8\r\n" +
			jobs.Body + "\r\n",
	)

	address := fmt.Sprintf("%s:%s", host, port)

	emailPayload := &emails.IPayload{
		User_id:  0,
		Email_id: strconv.Itoa(jobs.ID),
		Sender:   jobs.Sender,
		Receiver: jobs.To,
		Subject:  jobs.Subject,
		Body:     jobs.Body,
	}
	go func() {
		_, err := emails.EmailSerice(db).CreateNewEmail(context.Background(), *emailPayload)
		if err != nil {
			log.Fatal(err)
		}

	}()
	return smtp.SendMail(address, auth, username, []string{jobs.To}, message)
}

func SMTPFunction(jobs model.IEmailJob) {

}

// Resend Provedir
type ResendStruct struct {
	client *resend.Client
	Ctx    context.Context
	from   string
}

func NewResend(ctx context.Context, apiKey string, from string) *ResendStruct {
	return &ResendStruct{
		client: resend.NewClient(apiKey),
		Ctx:    ctx,
		from:   from,
	}
}

func (r *ResendStruct) ResendFunction(jobs model.IEmailJob) error {
	if r.client == nil {
		return errors.New("There error in api key")
	}

	params := &resend.SendEmailRequest{
		From:    r.from,
		To:      []string{jobs.To},
		Subject: jobs.Sender,
		Html:    jobs.Sender,
	}

	_, err := r.client.Emails.SendWithContext(r.Ctx, params)
	if err != nil {
		return err
	}

	return nil
}
