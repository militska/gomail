package internal

import (
	"net/smtp"
	"os"
)

type SMTPServer struct {
	Auth    smtp.Auth
	Host    string
	Port    int
	To      []string
	From    string
	Message []byte
}

func (a *SMTPServer) GetAuthData() smtp.Auth {
	data := smtp.PlainAuth("",
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("HOST_SMTP"))
	return data
}
