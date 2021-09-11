package internal

import "errors"

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

func (m *Msg) GetText() ([]byte, error) {

	if err := m.Check(); err != nil {
		return nil, err
	}

	message := []byte(
		"To: " + m.To + "\r\n" +
			"From: " + m.From + "\r\n" +
			"Subject: " + m.Subject + "\r\n" +
			"\r\n" + m.Body + "\r\n")

	return message, nil
}
