package handlers

import (
	"fmt"
	"net/smtp"
)

func (a *NotificationsService) A() {
	fmt.Println("nose")
}

// handles internally all the required logic to send an email
// by given values, it can send html-compliant mails
func SendEmail(to []string, topic, message string) error {
	user := ""
	password := ""
	host := ""
	port := ""

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
		return err
	}

	return err
}
