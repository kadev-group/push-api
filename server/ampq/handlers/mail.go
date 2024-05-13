package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"push-api/internal/interfaces"
	"push-api/internal/models"
)

type MailsHandler struct {
	log     *zap.Logger
	config  *models.Config
	manager interfaces.IManager
}

func InitMailsHandler(
	config *models.Config,
	manager interfaces.IManager,
	log *zap.Logger) (*MailsHandler, error) {
	mh := &MailsHandler{
		log:     log,
		config:  config,
		manager: manager,
	}
	if err := mh.Consume(); err != nil {
		return nil, err
	}
	return mh, nil
}

func (mh *MailsHandler) Consume() error {
	ctx := context.Background()
	ch, err := mh.manager.Processor().Queue().Consumers().Mails().Consume()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case body, ok := <-ch:
				if !ok {
					// TODO reconnect
					mh.log.Fatal("something went wrong")
				}
				mh.handleMessage(ctx, body)
			default:
			}
		}
	}()
	return nil
}

func (mh *MailsHandler) handleMessage(_ context.Context, body []byte) {
	log := mh.log.With(zap.String("body", string(body))).
		Named("handleMessage")

	msg := &models.MailsConsumerMsg{}

	err := json.Unmarshal(body, msg)
	if err != nil {
		log.Error(fmt.Sprintf("unmarshal: %s", err))
		return
	}

	if err = mh.manager.Service().Mail().SendVerificationCode(msg.SendTo, msg.VerificationCode); err != nil {
		log.Error(fmt.Sprintf("s.SendVerificationCode: %s", err))
	}

	log.Info("message sent")
}
