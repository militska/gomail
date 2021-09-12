package internal

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/smtp"
	"os"
)

type Email struct {
	From         string
	FromExtended string
	To           []string
	Subject      string
	Body         string
}

func (vo *Email) Check() error {
	if vo.Subject == "" {
		return errors.New("Subject is empty")
	}
	if vo.Body == "" {
		return errors.New("Body is empty")
	}
	if vo.From == "" {
		return errors.New("From is empty")
	}
	if len(vo.To) == 0 {
		return errors.New("To is empty")
	}

	return nil
}

func SendMsg(data *Email) {
	msg := Msg{
		From:    data.FromExtended,
		Subject: data.Subject,
		Body:    data.Body,
	}

	err := msg.Check()
	if err != nil {
		log.Warning(err)
		return
	}

	server := SmtpServer{
		To:      data.To,
		From:    data.From,
		Message: msg.GetText(),
	}
	err = smtp.SendMail(
		fmt.Sprintf("%s:%s", os.Getenv("HOST_SMTP"), os.Getenv("PORT_SMTP")),
		server.GetAuthData(),
		server.From,
		server.To,
		server.Message,
	)

	if err != nil {
		log.Fatal(err)
	}
}
