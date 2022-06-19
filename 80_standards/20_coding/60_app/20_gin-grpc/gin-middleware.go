package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("[GIN Server] LoggerMiddleware: start at ", time.Now())

		c.Next()

		log.Println("[GIN Server] LoggerMiddleware: duration is ", time.Since(t))
	}
}
