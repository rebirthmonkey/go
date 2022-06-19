package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginEngine := gin.Default()

	ginEngine.POST("/form", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
		})
	})

	ginEngine.Run(":8080")
}
