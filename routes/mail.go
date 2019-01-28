package routes

import (
	"net/http"
	"os"
	"time"

	sendgrid "github.com/sendgrid/sendgrid-go"
)

// Mail is this
type Mail struct {
	Receiver []string  `json:"receiver"`
	Message  string    `json:"message"`
	Subject  string    `json:"subject"`
	Sent     bool      `json:"sent"`
	Date     time.Time `json:"date"`
}

// MailHandler handles mail requests
// /v1/mail/send + body
func MailHandler(w http.ResponseWriter, r *http.Request) {
	sgclient := sendgrid.NewSendClient(os.Getenv("SENDGRID_API"))
	_ = sgclient
	// from := mail.NewEmail("from", "asdasd")
	// subject := "Testing the mail"
	// to := mail.NewEmail("Me asd ads", "simon@byrd.news")
	// mail := mail.NewSingleEmail(from, subject, to, "Hello! I am doing this!", "<a>CLick meeee</a>")
	// response, err := sgclient.Send(mail)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	c.JSON(201, gin.H{
	// 		"status":   "sent",
	// 		"receiver": response.StatusCode,
	// 		"message":  response.Body,
	// 		"subject":  response.Headers,
	// 	})
	// }
}
