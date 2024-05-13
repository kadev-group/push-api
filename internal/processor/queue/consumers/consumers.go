package consumers

import (
	"go.uber.org/zap"
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"sync"
)

type Consumers struct {
	log              *zap.Logger
	config           *models.Config
	consumerProvider interfaces.IQueueConsumerProvider

	mailsConsumer       interfaces.IQueueConsumerProcessor
	mailsConsumerRunner sync.Once
}

func NewConsumersProcessor(
	config *models.Config,
	consumersProvider interfaces.IQueueConsumerProvider,
	log *zap.Logger) *Consumers {
	return &Consumers{
		log:              log,
		config:           config,
		consumerProvider: consumersProvider,
	}
}

func (c *Consumers) Mails() interfaces.IQueueConsumerProcessor {
	c.mailsConsumerRunner.Do(func() {
		c.mailsConsumer = NewMailsConsumer(c.config, c.consumerProvider, c.log.Named("[MAILs]"))
	})
	return c.mailsConsumer
}
