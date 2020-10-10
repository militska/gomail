package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func send_msg() {
	authStr := AuthStr{}
	auth := authStr.getData()

	to := []string{"cheshirenok@gmail.com"}
	from := "militska.ru@gmail.com"
	msg := Msg{
		From: from,
		//From:    "test change <militska.ru@gmail.com>",
		To:      "",
		Subject: "test title 333",
		Body:    "test body 222",
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
