package services

import (
	"os"

	"gopkg.in/gomail.v2"
)

var Mailer *MailService

type MailService struct {
	dialer *gomail.Dialer
	from   string
}

func NewMailService(host string, port int, username, password, from string) *MailService {
	d := gomail.NewDialer(host, port, username, password)
	return &MailService{
		dialer: d,
		from:   from,
	}
}

func (m *MailService) SendMail(to, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	return m.dialer.DialAndSend(msg)
}

func InitiateMailService() {
	Mailer = NewMailService(
		"smtp.gmail.com",
		587,
		os.Getenv("MAIL"),
		os.Getenv("PASSWORD"),
		os.Getenv("MAIL"),
	)
}
