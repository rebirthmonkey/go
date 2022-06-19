package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "12345")

		c.Next()

		latency := time.Since(t)
		log.Println("latency is:", latency)

		status := c.Writer.Status()
		log.Println("status is:", status)
	}
}

func main() {
	ginEngine := gin.Default()
	ginEngine.Use(Logger())

	ginEngine.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println("GET /test, example is:", example) // it would print: "12345"
	})

	ginEngine.Run(":8080")
}
