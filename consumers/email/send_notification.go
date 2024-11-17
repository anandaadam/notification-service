package email

import (
	"fmt"
	"notification-service/config"
	"notification-service/helpers"

	"gopkg.in/gomail.v2"
)

func SendNotification(messageBody *helpers.MessageBody) error {
	config := config.GetSMTPConfig()

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.EmailFrom)
	mailer.SetHeader("To", messageBody.SendTo)
	mailer.SetHeader("Subject", messageBody.Subject)
	mailer.SetBody("text/plain", messageBody.Message)

	dialer := gomail.NewDialer(config.Server, config.Port, config.Username, config.Password)

	if err := dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}
