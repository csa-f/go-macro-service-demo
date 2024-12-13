package repository

import (
	"github.com/csa-f/go-macro-service-demo/common/conf"
	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	Config *conf.Config
}
