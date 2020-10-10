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
	msg := []byte("To: cheshirenok@gmail.com\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail(os.Getenv("HOST")+":587", auth, os.Getenv("EMAIL_USER")+"@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
