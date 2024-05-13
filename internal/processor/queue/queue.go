package queue

import (
	"go.uber.org/zap"
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"push-api/internal/processor/queue/consumers"
	"sync"
)

type Queue struct {
	log              *zap.Logger
	config           *models.Config
	consumerProvider interfaces.IQueueConsumerProvider

	consumersProcessor       interfaces.IQueueConsumersProcessor
	consumersProcessorRunner sync.Once
}

func NewQueueProcessor(
	config *models.Config,
	consumersProvider interfaces.IQueueConsumerProvider,
	log *zap.Logger) *Queue {
	return &Queue{
		log:              log,
		config:           config,
		consumerProvider: consumersProvider,
	}
}

func (q *Queue) Consumers() interfaces.IQueueConsumersProcessor {
	q.consumersProcessorRunner.Do(func() {
		q.consumersProcessor = consumers.NewConsumersProcessor(q.config, q.consumerProvider, q.log.Named("[CONSUMERS]"))
	})
	return q.consumersProcessor
}
