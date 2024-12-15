package repository

import (
	"github.com/csa-f/go-macro-service-demo/middle/config"
	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewRepository(
	db *gorm.DB,
	c *config.Config,
) *Repository {
	return &Repository{
		db,
		c,
	}
}
