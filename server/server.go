package server

import (
	"go.uber.org/zap"
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"push-api/server/ampq"
	"push-api/server/rest"
	"sync"
)

type Server struct {
	log     *zap.Logger
	config  *models.Config
	manager interfaces.IManager

	restServer       interfaces.IRESTServer
	restServerRunner sync.Once

	ampqServer       interfaces.IAMPQServer
	ampqServerRunner sync.Once
}

func InitServer(
	config *models.Config,
	manager interfaces.IManager,
	log *zap.Logger) *Server {
	return &Server{
		log:     log,
		config:  config,
		manager: manager,
	}
}

func (p *Server) REST() interfaces.IRESTServer {
	p.restServerRunner.Do(func() {
		p.restServer = rest.InitREST(p.config, p.manager.Service(), p.log.Named("[REST]"))
	})
	return p.restServer
}

func (p *Server) AMPQ() interfaces.IAMPQServer {
	p.ampqServerRunner.Do(func() {
		p.ampqServer = ampq.InitAMPQ(p.config, p.manager, p.log.Named("[AMPQ]"))
	})
	return p.ampqServer
}
