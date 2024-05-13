package smtp

import (
	"bytes"
	"github.com/doxanocap/pkg/errs"
	"gopkg.in/gomail.v2"
	"log"
	"push-api/internal/models"
)

const (
	ContentTypeHTML = "text/html"
)

type Mailer struct {
	config *models.Config
	gomail *gomail.Dialer
}

func NewSMTPMailer(config *models.Config) *Mailer {
	dialer := gomail.NewDialer(
		config.SMTP.SmtpHost,
		config.SMTP.SmtpPort,
		config.SMTP.SmtpUsername,
		config.SMTP.SmtpPassword)
	_, err := dialer.Dial()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Mailer{
		config: config,
		gomail: dialer,
	}
}

type MailMessage struct {
	Subject string
	SendTo  []string

	ContentType string
	Body        *bytes.Buffer

	Attachments []string
}

func (m *Mailer) Send(mailMsg *MailMessage) error {
	message := gomail.NewMessage()

	message.SetHeader("From", m.config.SMTP.SmtpUsername)

	if mailMsg.Subject != "" {
		message.SetHeader("Subject", mailMsg.Subject)
	}
	if !isValidContentType(mailMsg.ContentType) {
		return errs.New("invalid content type")
	}

	if mailMsg.Body != nil {
		message.SetBody(mailMsg.ContentType, mailMsg.Body.String())
	}

	message.SetHeader("To", mailMsg.SendTo...)
	for _, f := range mailMsg.Attachments {
		message.Attach(f)
	}

	if err := m.gomail.DialAndSend(message); err != nil {
		return err
	}
	return nil
}

func isValidContentType(ct string) bool {
	return ct == ContentTypeHTML
}
