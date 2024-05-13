package service

import (
	"push-api/internal/interfaces"
)

type MailerService struct {
	manager interfaces.IManager
}

func InitMailerService(manager interfaces.IManager) *MailerService {
	return &MailerService{
		manager: manager,
	}
}

func (m *MailerService) SendVerificationCode(email, code string) error {
	if err := m.manager.Processor().
		Mailer().
		SendVerificationCode(email, code); err != nil {
		return err
	}

	return nil
}
