package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.Default()
	g.Run(":8080")
}
