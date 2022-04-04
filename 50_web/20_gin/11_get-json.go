package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(context *gin.Context) {
	println(">>>> ping function start <<<<")

	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"message": "pong",
		"success":true,
	})
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.Run(":8080")
}

