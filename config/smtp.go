package config

import (
	"os"
	"strconv"
)

type SMTPConfig struct {
	Server    string
	Port      int
	Username  string
	Password  string
	EmailFrom string
}

func GetSMTPConfig() *SMTPConfig {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	return &SMTPConfig{
		Server:    os.Getenv("SMTP_SERVER"),
		Port:      port,
		Username:  os.Getenv("SMTP_USERNAME"),
		Password:  os.Getenv("SMTP_PASSWORD"),
		EmailFrom: os.Getenv("EMAIL_FROM"),
	}
}
