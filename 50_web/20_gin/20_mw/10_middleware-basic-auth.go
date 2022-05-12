package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginEngine := gin.Default()
	ginEngine.Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"}))

	ginEngine.GET("/auth", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Basic Auth")
	})

	ginEngine.Run(":8080")
}
