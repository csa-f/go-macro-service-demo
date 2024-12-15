package main

import (
	"fmt"

	BaseConf "github.com/csa-f/go-macro-service-demo/common/config"
	"github.com/gin-gonic/gin"
)

func main() {
	BaseConf.LoadConfig("gateway")
	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})
	server := BaseConf.GetConfig().Server
	if err := g.Run(fmt.Sprintf("%s:%s", server.Ip, server.Port)); err != nil {
		panic(fmt.Errorf("服务启动失败:%s", err))
	}
}
