package consumers

import (
	"go.uber.org/zap"
	"push-api/internal/interfaces"
	"push-api/internal/models"
)

type MailsConsumer struct {
	log      *zap.Logger
	config   *models.Config
	provider interfaces.IQueueConsumerProvider
}

func NewMailsConsumer(
	config *models.Config,
	provider interfaces.IQueueConsumerProvider,
	log *zap.Logger) *MailsConsumer {
	return &MailsConsumer{
		log:      log,
		config:   config,
		provider: provider,
	}
}

func (mc *MailsConsumer) Consume() (chan []byte, error) {
	ch, err := mc.provider.Consume(mc.config.MailsQueue, "")
	if err != nil {
		return nil, err
	}
	return ch, nil
}
