package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	if err := g.Run(":8080"); err != nil {
		panic(fmt.Errorf("服务启动失败:%s", err))
	}
}
