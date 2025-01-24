package repository

import (
	"github.com/csa-f/go-macro-service-demo/middle/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	redis  *redis.Client
	Config *config.Config
}

func NewRepository(
	db *gorm.DB,
	redis *redis.Client,
	conf *config.Config,
) *Repository {
	return &Repository{
		DB:     db,
		redis:  redis,
		Config: conf,
	}
}
