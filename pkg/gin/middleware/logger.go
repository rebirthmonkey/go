package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/log"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.L(c).Infof("[Logger Middleware] Start at: %v", time.Now())

		c.Next()

		log.L(c).Infof("[Logger Middleware] duration is: %v", time.Since(t))
	}
}
