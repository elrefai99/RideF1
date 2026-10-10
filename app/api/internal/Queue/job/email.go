package JobQueue

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/elrefai99/RideF1/app/api/internal/Queue/model"
)

func SendEmail(jobs model.IEmailJob) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", username, password, host)

	message := []byte(
		"From: " + username + "\r\n" +
			"To: " + jobs.To + "\r\n" +
			"Subject: Hi" + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain;charset=UTF-8\r\n" +
			jobs.Body + "\r\n",
	)

	address := fmt.Sprintf("%s:%s", host, port)

	return smtp.SendMail(address, auth, username, []string{jobs.To}, message)
}
