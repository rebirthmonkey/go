package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware3")
		c.Set("request3", "AAA")
		c.Set("request4", "AAA")
		c.Next()
		fmt.Println("after middleware3")
	}
}

func MiddleWare4() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware4")
		c.Set("request3", "BBB")
		c.Next()
		fmt.Println("after middleware4")
	}
}

func main(){
	router := gin.Default()

	router.GET("/before", MiddleWare3(), func(c *gin.Context) {
		request3 := c.MustGet("request3").(string)
		request4 := c.MustGet("request4").(string)
		c.JSON(http.StatusOK, gin.H{
			"request3": request3,
			"request4": request4,
		})
	})

	router.GET("/after", MiddleWare4(), func(c *gin.Context) {
		request3 := c.MustGet("request3").(string)
		request4 := c.MustGet("request4").(string)
		c.JSON(http.StatusOK, gin.H{
			"request3": request3,
			"request4": request4,
		})
	})

	router.Run(":8080")
}
