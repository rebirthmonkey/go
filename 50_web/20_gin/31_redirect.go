package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()

	router.GET("/redirect/google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://google.com")
	})

	router.Run(":8080")
}

