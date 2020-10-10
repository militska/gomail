package main

import "go/types"

type Msg struct {
	To      string
	From    string
	Subject string
	Body    string
}

func (m *Msg) isFill() error {
	if m.Body == "" {
		return types.Error{Msg: "Body is empty"}
	}
	return nil
}

func (m *Msg) getText() ([]byte, error) {

	if err := m.isFill(); err != nil {
		return nil, err
	}

	message := []byte(
		"To: " + m.To + "\r\n" +
			"From: " + m.From + "\r\n" +
			"Subject: " + m.Subject + "\r\n" +
			"\r\n" + m.Body + "\r\n")

	return message, nil
}
