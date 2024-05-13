package mailer

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"html/template"
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"push-api/internal/pkg/smtp"
)

const (
	verificationCodeKey = "verification_code"
)

type Mailer struct {
	log       *zap.Logger
	config    *models.Config
	templates map[string]*template.Template
	provider  interfaces.IMailerProvider
}

func NewMailerProcessor(config *models.Config, provider interfaces.IMailerProvider, log *zap.Logger) *Mailer {
	templates := map[string]*template.Template{}

	tmpl, err := template.ParseFiles("/web/verification_code.html")
	if err != nil {
		log.Fatal(fmt.Sprintf("parse template: %s", err))
	}
	templates[verificationCodeKey] = tmpl

	return &Mailer{
		log:       log,
		config:    config,
		templates: templates,
		provider:  provider,
	}
}

type verifyCodeData struct {
	VerificationCode string
}

func (m *Mailer) SendVerificationCode(sendTo, code string) error {
	log := m.log.With(zap.String("send_to", sendTo),
		zap.String("code", code)).
		Named("SendVerificationCode")

	data := verifyCodeData{VerificationCode: code}
	body := &bytes.Buffer{}
	if err := m.templates[verificationCodeKey].Execute(body, data); err != nil {
		return err
	}

	go func() {
		err := m.provider.Send(&smtp.MailMessage{
			Subject:     "Kantoo: verification code",
			SendTo:      []string{sendTo},
			ContentType: smtp.ContentTypeHTML,
			Body:        body,
		})
		if err != nil {
			log.Error(fmt.Sprintf("provider.Send: %s", err))
		}
	}()

	log.Info("success")
	return nil
}
