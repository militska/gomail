package main

import (
	"encoding/json"
	//"encoding/json"
	"io/ioutil"

	//"fmt"
	"github.com/joho/godotenv"
	//"html"
	"log"
	"net/http"
	"os"
	"time"
)

const ENABLED = "1"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	if os.Getenv("ENABLED_API") == ENABLED {
		s := &http.Server{
			Addr:           ":8070",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			var t SendEmailDTO
			err = json.Unmarshal(body, &t)
			if err != nil {
				panic(err)
			}
			log.Println(t.Body)

			send_msg(t)
		})

		log.Fatal(s.ListenAndServe())
	}

}

type SendEmailDTO struct {
	From         string
	FromExtended string
	To           []string
	Subject      string
	Body         string
}

func send_msg(dto SendEmailDTO) {
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
