package util

import (
	"fmt"
	"net/smtp"

	"server/env"

	"github.com/jordan-wright/email"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (s *GmailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attacthFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.name, s.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	smtpAuth := smtp.PlainAuth("", s.fromEmailAddress, s.fromEmailPassword, env.NewEnvironment().SMTP_HOST)
	return e.Send(env.NewEnvironment().SMTP_HOST+":"+env.NewEnvironment().SMTP_PORT, smtpAuth)
}
