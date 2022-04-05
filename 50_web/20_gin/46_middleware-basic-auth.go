package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.New()
	r.Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"}))


	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Basic Auth")
	})

	r.Run(":8080")
}