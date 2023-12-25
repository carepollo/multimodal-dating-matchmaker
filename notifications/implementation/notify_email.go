package implementation

import (
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/carepollo/multimodal-dating-matchmaker/protos"
)

// NotifyByEmail handles internally all the required logic to send an email
// by given values, it can send html-compliant mails
func (service *NotificationsService) NotifyByEmail(ctx context.Context, req *protos.NotifyEmailRequest) (*protos.NotifyEmailResponse, error) {
	user := os.Getenv("EMAIL_USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	topic := req.Topic
	message := req.Message
	to := req.To

	auth := smtp.PlainAuth(
		"",
		user,
		password,
		host,
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	body := fmt.Sprintf("Subject: %v\n%v\n\n%v", topic, headers, message)
	address := fmt.Sprintf("%v:%v", host, port)

	err := smtp.SendMail(address, auth, "system@librecode.com", to, []byte(body))
	if err != nil {
		return &protos.NotifyEmailResponse{Message: err.Error()}, err
	}

	return &protos.NotifyEmailResponse{Message: "email notification successfully sent"}, err
}
