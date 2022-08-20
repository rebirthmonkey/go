package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("Start at:", time.Now())

		c.Next()

		log.Println("duration is:", time.Since(t))
	}
}
