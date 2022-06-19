package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/pingHandler", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // 输出json结果给调用方
			"message": "pong",
		})
	})
	ginEngine.Run(":8080")
}
