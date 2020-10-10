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

type SendEmailDTO struct {
	From         string
	FromExtended string
	To           []string
	Subject      string
	Body         string
}

type AuthStr struct{}

func sendMsg(dto SendEmailDTO) {
	authStr := AuthStr{}
	auth := authStr.getData()

	to := dto.To
	from := dto.From
	msg := Msg{
		From:    dto.FromExtended,
		Subject: dto.Subject,
		Body:    dto.Body,
	}

	text, _ := msg.getText()

	server := SmtpServer{
		Auth:    auth,
		To:      to,
		From:    from,
		Message: text,
	}
	server.send()
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

func (a *AuthStr) getData() smtp.Auth {
	data := smtp.PlainAuth("",
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("HOST"))
	return data
}
