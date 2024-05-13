package main

import (
	"github.com/doxanocap/pkg/config"
	"github.com/doxanocap/pkg/logger"
	"go.uber.org/fx"
	"log"
	"push-api/internal/manager"
	"push-api/internal/models"
	"push-api/internal/pkg/postgres"
	"push-api/internal/pkg/rabbitmq"
	"push-api/internal/pkg/redis"
	"push-api/internal/pkg/smtp"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.InitConfig[models.Config],
			logger.InitLogger[models.Config],
			smtp.NewSMTPMailer,
			rabbitmq.NewConsumerClient,
			postgres.InitConnection,
			redis.InitConnection,
			manager.InitManager,
		),
		fx.Invoke(manager.Run),
	)

	app.Run()
	if err := app.Err(); err != nil {
		log.Fatal(err)
	}
}
