package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{  // 输出json结果给调用方
			"message": "pong",
		})
	})
	router.Run(":8080")
}