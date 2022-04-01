package email

import (
	"bytes"
	"errors"
	"html/template"
	"net/mail"
)

var (
	ErrEmailEmptyTo      = errors.New("to cannot be empty")
	ErrEmailEmptySubject = errors.New("subject cannot be empty")
	ErrEmailEmptyMessage = errors.New("message cannot be empty")
)

type Letter struct {
	To      string
	Subject string
	Message string
}

func (l *Letter) SetMsgFromTemplate(filename string, data interface{}) error {
	t, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, &data); err != nil {
		return err
	}

	l.Message = buf.String()

	return nil
}

func (l *Letter) Validate() error {
	if l.To == "" {
		return ErrEmailEmptyTo
	}

	_, err := mail.ParseAddress(l.To)
	if err != nil {
		return err
	}

	if l.Subject == "" {
		return ErrEmailEmptySubject
	}

	if l.Message == "" {
		return ErrEmailEmptyMessage
	}

	return nil
}
