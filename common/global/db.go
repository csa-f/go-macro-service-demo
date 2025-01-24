package global

import (
	"fmt"
	"log"

	"github.com/csa-f/go-macro-service-demo/common/config"
	"github.com/csa-f/go-macro-service-demo/common/consts"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
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
	case consts.DriverSqlite:
		dialector = sqlite.Open(dsn)
	case consts.DriverPostgres:
		dialector = postgres.Open(dsn)
	case consts.DriverSqlserver:
		dialector = sqlserver.Open(dsn)
	case consts.DriverClickhouse:
		dialector = clickhouse.Open(dsn)
	default:
		log.Fatalf("db driver not support:%s", conf.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("db connect fail:%s", err)
	}
	return db
}
