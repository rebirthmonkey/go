package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware3")
		c.Set("request", "AAA")
		c.Next()
		fmt.Println("after middleware3")
	}
}

func MiddleWare4() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware4")
		c.Set("request", "BBB")
		c.Next()
		fmt.Println("after middleware4")
	}
}

func main(){
	router := gin.Default()

	router.GET("/before", MiddleWare3(), func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"request": request,
		})
	})

	router.GET("/after", MiddleWare4(), func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"request": request,
		})
	})

	router.Run(":8080")
}
