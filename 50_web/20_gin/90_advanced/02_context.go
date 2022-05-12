package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/", func(c *gin.Context) {
		c.Set("xxx", "1111")
		c.Set("yyy", "2222")
		fmt.Println("context key-values are:", c.Keys)
		fmt.Println("context Request headers are:", c.Request.Header)
		c.String(http.StatusOK, "Hello World")
		fmt.Println("context Response headers are:", c.Writer.Header())
	})
	ginEngine.Run(":8080")
}
