package service

import (
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"sync"
)

type Service struct {
	config  *models.Config
	manager interfaces.IManager

	mail       interfaces.IMailService
	mailRunner sync.Once
}

func InitService(manager interfaces.IManager, config *models.Config) *Service {
	return &Service{
		config:  config,
		manager: manager,
	}
}

func (s *Service) Mail() interfaces.IMailService {
	s.mailRunner.Do(func() {
		s.mail = InitMailerService(s.manager)
	})
	return s.mail
}
