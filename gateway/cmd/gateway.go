package main

import (
	"fmt"

	CommonConfig "github.com/csa-f/go-macro-service-demo/common/config"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})
	server := CommonConfig.Get("app").Server
	if err := g.Run(fmt.Sprintf("%s:%d", server.Ip, server.Port)); err != nil {
		panic(fmt.Errorf("服务启动失败:%s", err))
	}
}
