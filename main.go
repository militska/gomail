package main

import (
	log "github.com/sirupsen/logrus"
	"militska/gomail/cmd"
	"militska/gomail/config"
)

func main() {
	err := config.LoadEnv()

	if err != nil {
		log.Warning(err)
	}

	cmd.InitHTTPServer()
}
