package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置一个新变量 example
		c.Set("example", "12345")

		// 请求之前
		c.Next()

		// 请求之后
		latency := time.Since(t)
		log.Println("latency is:", latency)

		// 发送状态
		status := c.Writer.Status()
		log.Println("status is:", status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println("GET /test, example is:", example)	// it would print: "12345"
	})

	r.Run(":8080")
}