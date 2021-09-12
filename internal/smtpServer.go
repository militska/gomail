package internal

import (
	"net/smtp"
	"os"
)

type SmtpServer struct {
	Auth    smtp.Auth
	Host    string
	Port    int
	To      []string
	From    string
	Message []byte
}

func (a *SmtpServer) GetAuthData() smtp.Auth {
	data := smtp.PlainAuth("",
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("HOST_SMTP"))
	return data
}
