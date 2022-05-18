package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Gin")
	})
	service := router.Group("/service")
	service.GET("/list", func(ctx *gin.Context) {
		ctx.String(200, "list")
	})
	router.Run(":8080")
}
