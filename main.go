package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	auth := smtp.PlainAuth("", os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("HOST"))

	to := []string{"cheshirenok@gmail.com"}

	msg := Msg{From: "test change <militska.ru@gmail.com>", To: "", Subject: "test title", Body: "test body"}

	err := smtp.SendMail(os.Getenv("HOST")+":587", auth, os.Getenv("EMAIL_USER")+"@gmail.com", to, msg.getText())
	if err != nil {
		log.Fatal(err)
	}
}

type Msg struct {
	To      string
	From    string
	Subject string
	Body    string
}

func (m *Msg) getText() []byte {

	message := []byte(
		"To: " + m.To + "\r\n" +
			"From: " + m.From + "\r\n" +
			"Subject: " + m.Subject + "\r\n" +
			"\r\n" + m.Body + "\r\n")

	return message
}
