package main

import (
	"github.com/csa-f/go-macro-service-demo/common/global"
	"github.com/csa-f/go-macro-service-demo/middle/cmd/server"
	"github.com/csa-f/go-macro-service-demo/middle/config"
)

func main() {
	appConf := config.Get()
	global.InitLog(appConf.Log)
	db := global.InitDB(appConf.DB)
	redis := global.InitRedis(appConf.Redis)
	server.NewRpcServer(db, redis, appConf)
}
