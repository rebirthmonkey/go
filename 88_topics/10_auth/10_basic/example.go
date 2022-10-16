package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.New()

	authStrategy := NewBasicStrategy(func(username string, password string) bool {
		if username == "admin" && password == "admin" {
			return true
		}
		return false
	})
	ginEngine.Use(authStrategy.AuthFunc())

	ginEngine.GET("/ping", func(c *gin.Context) {
		usernameKey := c.MustGet(UsernameKey).(string)
		c.JSON(http.StatusOK, gin.H{
			"UsernameKey": usernameKey,
			"message":     "pong",
		})
	})

	ginEngine.Run(":8080")
}
