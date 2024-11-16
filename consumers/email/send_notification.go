package email

import (
	"fmt"
	"notification-service/helpers"

	"gopkg.in/gomail.v2"
)

func SendNotification(messageBody *helpers.MessageBody) error {
	smtpServer := "smtp.mailtrap.io"
	smtpPort := 587
	username := "3d8689177baa73"
	password := "e8d5b794377601"

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "adamanandasantoso@gmail.com")
	mailer.SetHeader("To", messageBody.SendTo)
	mailer.SetHeader("Subject", "Success")
	mailer.SetBody("text/plain", messageBody.Message)

	dialer := gomail.NewDialer(smtpServer, smtpPort, username, password)

	if err := dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}
