package manager

import (
	_ "github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"push-api/internal/pkg/rabbitmq"
	"push-api/internal/pkg/redis"
	"push-api/internal/pkg/smtp"
	"push-api/internal/processor"
	"push-api/internal/repository"
	"push-api/internal/service"
	"push-api/server"
	"sync"
)

type Manager struct {
	db                    *sqlx.DB
	log                   *zap.Logger
	config                *models.Config
	mailerProvider        *smtp.Mailer
	cacheProvider         interfaces.ICacheProvider
	queueConsumerProvider interfaces.IQueueConsumerProvider
	service               interfaces.IService
	serviceRunner         sync.Once

	repository       interfaces.IRepository
	repositoryRunner sync.Once

	processor       interfaces.IProcessor
	processorRunner sync.Once

	server       interfaces.IServer
	serverRunner sync.Once
}

func InitManager(
	db *sqlx.DB,
	log *zap.Logger,
	config *models.Config,
	mailer *smtp.Mailer,
	rabbitmqConsumer *rabbitmq.ConsumerClient,
	redisConn *redis.Conn) *Manager {
	return &Manager{
		db:                    db,
		log:                   log,
		config:                config,
		cacheProvider:         redisConn,
		mailerProvider:        mailer,
		queueConsumerProvider: rabbitmqConsumer,
	}
}

func (m *Manager) Repository() interfaces.IRepository {
	m.repositoryRunner.Do(func() {
		m.repository = repository.InitRepository(m.db, m.config, m.Processor().Cache(), m.log.Named("[REPOSITORY]"))
	})
	return m.repository
}

func (m *Manager) Service() interfaces.IService {
	m.serviceRunner.Do(func() {
		m.service = service.InitService(m, m.config)
	})
	return m.service
}

func (m *Manager) Processor() interfaces.IProcessor {
	m.processorRunner.Do(func() {
		m.processor = processor.InitProcessor(m.log, m.config,
			m.Service(), m.cacheProvider, m.queueConsumerProvider, m.mailerProvider)
	})
	return m.processor
}

func (m *Manager) Server() interfaces.IServer {
	m.serverRunner.Do(func() {
		m.server = server.InitServer(m.config, m, m.log)
	})
	return m.server
}

func (m *Manager) SetDB(db *sqlx.DB) {
	m.db = db
}

func (m *Manager) SetCache(cache interfaces.ICacheProvider) {
	m.cacheProvider = cache
}
