package main

import (
	"errors"
	"github.com/joho/godotenv"
	"go/types"
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

	if errCheckEnv := checkEnvVariables(); errCheckEnv != nil {
		log.Print(errCheckEnv.Error())
	}

	initHttpServer()
}

func checkEnvVariables() error {
	if os.Getenv("HOST_SMTP") == "" {
		return errors.New("HOST_SMTP is empty")
	}
	if os.Getenv("PORT_SMTP") == "" {
		return errors.New("ENABLED_API is empty")
	}
	if os.Getenv("EMAIL_USER") == "" {
		return errors.New("ENABLED_API is empty")
	}
	if os.Getenv("EMAIL_PASSWORD") == "" {
		return errors.New("ENABLED_API is empty")
	}
	if os.Getenv("ENABLED_API") == "" {
		return types.Error{Msg: "ENABLED_API is empty"}
	}

	return nil
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
