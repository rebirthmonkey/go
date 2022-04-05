package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware1")
		c.Set("request1", "XXX")
		c.Set("request2", "XXX")
		c.Next()
		fmt.Println("after middleware1")
	}
}

func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware2")
		c.Set("request1", "YYY")
		c.Next()
		fmt.Println("after middleware2")
	}
}

func main(){
	router := gin.Default()

	router.Use(MiddleWare1())

	router.GET("/before", func(c *gin.Context) {
		request1 := c.MustGet("request1").(string)
		request2 := c.MustGet("request2").(string)
		c.JSON(http.StatusOK, gin.H{
			"request1": request1,
			"request2": request2,
		})
	})

	router.Use(MiddleWare2())

	router.GET("/after", func(c *gin.Context) {
		request1 := c.MustGet("request1").(string)
		request2 := c.MustGet("request2").(string)
		c.JSON(http.StatusOK, gin.H{
			"request1": request1,
			"request2": request2,
		})
	})

	router.Run(":8080")
}
