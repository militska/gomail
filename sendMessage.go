package main

import (
	"errors"
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

type SendEmailVo struct {
	From         string
	FromExtended string
	To           []string
	Subject      string
	Body         string
}

func (vo *SendEmailVo) check() error {
	if vo.Subject == "" {
		return errors.New("Subject is empty")
	}
	if vo.Body == "" {
		return errors.New("Body is empty")
	}
	if vo.From == "" {
		return errors.New("From is empty")
	}
	if len(vo.To) == 0 {
		return errors.New("To is empty")
	}

	return nil
}

func sendMsg(data SendEmailVo) {
	to := data.To
	from := data.From
	msg := Msg{
		From:    data.FromExtended,
		Subject: data.Subject,
		Body:    data.Body,
	}

	text, _ := msg.getText()

	server := SmtpServer{
		To:      to,
		From:    from,
		Message: text,
	}

	server.send()
}

func (server *SmtpServer) send() {
	err := smtp.SendMail(os.Getenv("HOST_SMTP")+":"+os.Getenv("PORT_SMTP"),
		server.getAuthData(),
		server.From,
		server.To,
		server.Message,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func (a *SmtpServer) getAuthData() smtp.Auth {
	data := smtp.PlainAuth("",
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("HOST_SMTP"))
	return data
}
