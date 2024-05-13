package interfaces

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"push-api/internal/pkg/smtp"
	"time"
)

type IProcessor interface {
	Cache() ICacheProcessor
	Mailer() IMailerProcessor
	Queue() IQueueProcessor
}

// Cache

type ICacheProcessor interface {
	Set(ctx context.Context, key string, value []byte) error
	SetJSON(ctx context.Context, key string, value interface{}) error
	SetJSONWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	GetJSON(ctx context.Context, key string, v interface{}) error
	Delete(ctx context.Context, key string) error
	FlushAll(ctx context.Context) error
}

type ICacheProvider interface {
	Set(ctx context.Context, key string, value []byte) error
	SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
	FlushAll(ctx context.Context) error
	Close() error
}

// Mailer

type IMailerProcessor interface {
	SendVerificationCode(sendTo, code string) error
}

type IMailerProvider interface {
	Send(mailMsg *smtp.MailMessage) error
}

// Queue

type IQueueProcessor interface {
	Consumers() IQueueConsumersProcessor
}

type IQueueConsumersProcessor interface {
	Mails() IQueueConsumerProcessor
}

type IQueueConsumerProcessor interface {
	Consume() (chan []byte, error)
}

type IQueueConsumerProvider interface {
	Consume(queueName string, consumer string, args ...amqp.Table) (chan []byte, error)
}

//
