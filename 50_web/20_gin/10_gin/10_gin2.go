package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func ping(context *gin.Context) {
	println(">>>> ping function start <<<<")

	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"message": "pong",
		"success":true,
	})
}