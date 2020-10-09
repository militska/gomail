package main

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gopkg.in/gomail.v2"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	to := "cheshirenok@gmail.com"

	d := gomail.NewDialer(os.Getenv("HOST"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "mewmew@mew.mew")
	m.SetHeader("To", to)
	m.SetBody("text/plain", "new text")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
