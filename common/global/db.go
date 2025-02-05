package global

import (
	"fmt"
	"log"

	"github.com/csa-f/macro-service/common/config"
	"github.com/csa-f/macro-service/common/consts"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(conf *config.DB) *gorm.DB {
	if conf.Host == "" {
		return nil
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.DBName,
		conf.Charset,
	)

	var dialector gorm.Dialector
	switch conf.Driver {
	case consts.DriverMysql:
		dialector = mysql.Open(dsn)
	case consts.DriverPostgres:
		dialector = postgres.Open(dsn)
	default:
		log.Fatalf("db driver not support:%s", conf.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("db connect fail:%s", err)
	}
	return db
}
