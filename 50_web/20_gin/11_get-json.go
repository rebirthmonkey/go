package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingHandler(context *gin.Context) {
	println(">>>> pingHandler function start <<<<")

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "pong",
		"success": true,
	})
}

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/ping", pingHandler)
	ginEngine.Run(":8080")
}
