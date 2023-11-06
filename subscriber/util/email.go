package util

import (
	"context"
	"fmt"
	"log"
	"subscriber/env"

	"github.com/hibiken/asynq"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

type emailStruct struct{}

func NewEmail() *emailStruct {
	return &emailStruct{}
}

type ResetPasswordStruct struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

func (e *emailStruct) ResetPassword(ctx context.Context, t *asynq.Task) error {
	a := UnmarshalConverter[ResetPasswordStruct](t.Payload())
	sender := NewGmailSender("Nama Aplikasi", env.ENV.SMTP.Email, env.ENV.SMTP.Password)
	subject := "Reset Password Verification"
	content := NewTemplate().EmailResetPassword(a.Token, a.ExpiredAt)
	to := []string{a.Email}
	if err := sender.SendEmail(subject, content, to, nil, nil, nil); err != nil {
		return fmt.Errorf("failed sent email to: %v", a.Email)
	}
	log.Printf("sending Email to: %v", a.Email)
	return nil
}

func (e *emailStruct) Register(ctx context.Context, t *asynq.Task) error {
	a := UnmarshalConverter[string](t.Payload())
	fmt.Println(a)
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", 1, "2")
	return nil
}
