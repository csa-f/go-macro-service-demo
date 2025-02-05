package main

import (
	"fmt"

	CommonConfig "github.com/csa-f/macro-service/common/config"
	"github.com/csa-f/macro-service/gateway/router"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	router.InitRouter(g)

	conf := CommonConfig.Get("app")
	server := conf.Server

	if err := g.Run(fmt.Sprintf("%s:%d", server.Ip, server.Port)); err != nil {
		panic(fmt.Errorf("服务启动失败:%s", err))
	}
}
