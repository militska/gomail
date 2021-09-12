package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

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
		return errors.New("ENABLED_API is empty")
	}

	return nil
}

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return errors.New("No .env file found")
	}

	if err := checkEnvVariables(); err != nil {
		return err
	}
	return nil
}
