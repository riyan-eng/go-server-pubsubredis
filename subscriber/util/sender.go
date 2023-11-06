package util

import (
	"fmt"
	"net/smtp"
	"subscriber/env"

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

type gmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &gmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (s *gmailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attacthFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.name, s.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	smtpAuth := smtp.PlainAuth("", s.fromEmailAddress, s.fromEmailPassword, env.ENV.SMTP.Host)
	smtpAdr := fmt.Sprintf("%v:%v", env.ENV.SMTP.Host, env.ENV.SMTP.Port)
	return e.Send(smtpAdr, smtpAuth)
}
