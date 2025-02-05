package router

import "github.com/gin-gonic/gin"

func InitRouter(g *gin.Engine) {
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})
}
