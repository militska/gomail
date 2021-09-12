package internal

import (
	"errors"
	"fmt"
)

type Msg struct {
	To      string
	From    string
	Subject string
	Body    string
}

func (m *Msg) Check() error {
	if m.Body == "" {
		return errors.New("Body is empty")
	}
	if m.Subject == "" {
		return errors.New("Subject is empty")
	}
	return nil
}

func (m *Msg) GetText() []byte {
	return []byte(
		fmt.Sprintf(
			`To: %s \r\n From: %s \r\n Subject: %s \r\n \r\n %s \r\n`,
			m.To, m.From, m.Subject, m.Body))

}
