package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	authStr := AuthStr{}
	auth := authStr.getData()
	to := []string{"cheshirenok@gmail.com"}

	msg := Msg{
		From:    "test change <militska.ru@gmail.com>",
		To:      "",
		Subject: "test title 333",
		Body:    "test body 222",
	}

	text, _ := msg.getText()

	server := SmtpServer{
		Auth:    auth,
		To:      to,
		From:    os.Getenv("EMAIL_USER") + "@gmail.com",
		Message: text,
	}
	server.send()
}

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
