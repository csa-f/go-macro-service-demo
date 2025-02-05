package global

import (
	"context"
	"fmt"
	"time"

	"github.com/csa-f/macro-service/common/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(conf *config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Username:     conf.User,
		Password:     conf.Pass,
		DB:           conf.DB,
		DialTimeout:  time.Millisecond * time.Duration(conf.DialTimeout),
		ReadTimeout:  time.Millisecond * time.Duration(conf.ReadTimeout),
		WriteTimeout: time.Millisecond * time.Duration(conf.WriteTimeout),
		PoolSize:     conf.PoolSize,
		PoolTimeout:  time.Millisecond * time.Duration(conf.PoolTimeout),
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic("redis connect failed [ERROR]=> " + err.Error())
	}
	return rdb
}
