package main

import (
	"log"
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

func (server *SmtpServer) send() {
	err := smtp.SendMail(os.Getenv("HOST")+":"+os.Getenv("PORT_SMTP"),
		server.Auth,
		server.From,
		server.To,
		server.Message,
	)

	if err != nil {
		log.Fatal(err)
	}
}

type AuthStr struct{}

func (a *AuthStr) getData() smtp.Auth {
	data := smtp.PlainAuth("",
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("HOST"))
	return data
}
