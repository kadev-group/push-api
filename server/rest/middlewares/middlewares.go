package middlewares

import (
	"go.uber.org/zap"
	"push-api/internal/interfaces"
)

type Middlewares struct {
	service interfaces.IService
	log     *zap.Logger
}

func InitMiddlewares(service interfaces.IService, log *zap.Logger) *Middlewares {
	return &Middlewares{
		service: service,
		log:     log,
	}
}
