package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"push-api/internal/interfaces"
	"push-api/internal/models"
)

type Repository struct {
	db     *sqlx.DB
	log    *zap.Logger
	config *models.Config
	cache  interfaces.ICacheProcessor
}

func InitRepository(
	db *sqlx.DB,
	config *models.Config,
	cache interfaces.ICacheProcessor,
	log *zap.Logger) *Repository {
	return &Repository{
		db:     db,
		log:    log,
		cache:  cache,
		config: config,
	}
}
