package processor

import (
	"go.uber.org/zap"
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"push-api/internal/processor/cache"
	"push-api/internal/processor/mailer"
	"push-api/internal/processor/queue"
	"sync"
)

type Processor struct {
	log                   *zap.Logger
	config                *models.Config
	service               interfaces.IService
	cacheProvider         interfaces.ICacheProvider
	mailerProvider        interfaces.IMailerProvider
	queueConsumerProvider interfaces.IQueueConsumerProvider

	cacheProcessor       interfaces.ICacheProcessor
	cacheProcessorRunner sync.Once

	mailerProcessor       interfaces.IMailerProcessor
	mailerProcessorRunner sync.Once

	queueProcessor       interfaces.IQueueProcessor
	queueProcessorRunner sync.Once
}

func InitProcessor(
	log *zap.Logger,
	config *models.Config,
	service interfaces.IService,
	cache interfaces.ICacheProvider,
	queueConsumerProvider interfaces.IQueueConsumerProvider,
	mailerProvider interfaces.IMailerProvider) *Processor {
	return &Processor{
		log:                   log,
		config:                config,
		service:               service,
		cacheProvider:         cache,
		mailerProvider:        mailerProvider,
		queueConsumerProvider: queueConsumerProvider,
	}
}

func (p *Processor) Cache() interfaces.ICacheProcessor {
	p.cacheProcessorRunner.Do(func() {
		p.cacheProcessor = cache.NewCacheProcessor(p.cacheProvider)
	})
	return p.cacheProcessor
}

func (p *Processor) Mailer() interfaces.IMailerProcessor {
	p.mailerProcessorRunner.Do(func() {
		p.mailerProcessor = mailer.NewMailerProcessor(p.config, p.mailerProvider, p.log.Named("[MAILER]"))
	})
	return p.mailerProcessor
}

func (p *Processor) Queue() interfaces.IQueueProcessor {
	p.queueProcessorRunner.Do(func() {
		p.queueProcessor = queue.NewQueueProcessor(p.config, p.queueConsumerProvider, p.log.Named("[QUEUE]"))
	})
	return p.queueProcessor
}
