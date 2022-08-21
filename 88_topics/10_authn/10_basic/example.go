package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	ginEngine.GET("/test", func(c *gin.Context) {
		usernameKey := c.MustGet(UsernameKey).(string)
		c.JSON(http.StatusOK, gin.H{
			"UsernameKey": usernameKey,
		})
		//c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	ginEngine.Run(":8080")
}
