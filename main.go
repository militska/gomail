package main

import (
	"github.com/joho/godotenv"
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

	initHttpServer()
}

func initHttpServer() {
	if os.Getenv("ENABLED_API") == ENABLED {
		s := &http.Server{
			Addr:           ":8070",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		sendMailHandler()

		log.Fatal(s.ListenAndServe())
	}
}
