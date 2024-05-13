package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"push-api/internal/models"
)

type ConsumerClient struct {
	chanel *amqp.Channel
}

func NewConsumerClient(config *models.Config, log *zap.Logger) *ConsumerClient {
	log = log.Named("[RabbitMQ]")

	conn, err := amqp.Dial(config.RabbitMQ.ServerURL)
	if err != nil {
		log.Fatal(fmt.Sprintf("connect: %s", err))
	}

	chanel, err := conn.Channel()
	if err != nil {
		log.Fatal(fmt.Sprintf("open channel: %s", err))
	}

	return &ConsumerClient{
		chanel: chanel,
	}
}

func (cc *ConsumerClient) Consume(queueName string, consumer string, args ...amqp.Table) (chan []byte, error) {
	var arg amqp.Table
	if len(args) != 0 {
		arg = args[0]
	}

	q, err := cc.chanel.QueueDeclare(queueName,
		true, false, false, false, arg)
	if err != nil {
		return nil, err
	}

	messages, err := cc.chanel.Consume(q.Name, consumer,
		false, false, false, false, arg)
	if err != nil {
		return nil, err
	}

	ch := make(chan []byte)
	go func() {
		for message := range messages {
			ch <- message.Body
		}
	}()
	return ch, nil
}
