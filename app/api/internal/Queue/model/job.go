package model

type IEmailJob struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	UserId  uint   `json:"userid"`
	Sender  string `json:"sender"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Payload any    `json:"payload"`
}
