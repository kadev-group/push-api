package ampq

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"push-api/server/ampq/handlers"
)

type AMPQ struct {
	log     *zap.Logger
	config  *models.Config
	manager interfaces.IManager

	mailsHandler *handlers.MailsHandler
}

func InitAMPQ(config *models.Config, manager interfaces.IManager, log *zap.Logger) *AMPQ {
	return &AMPQ{
		log:     log,
		config:  config,
		manager: manager,
	}
}

func (a *AMPQ) Handle() (err error) {
	a.mailsHandler, err = handlers.InitMailsHandler(a.config, a.manager, a.log.Named("[MAILs]"))
	if err != nil {
		log.Fatal(fmt.Sprintf("h.mailsConsumer: %s", err))
	}
	return nil
}
