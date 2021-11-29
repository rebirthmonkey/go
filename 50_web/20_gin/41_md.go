package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "XXX")
		c.Next()
		fmt.Println("before middleware")
	}
}

func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "YYY")
		c.Next()
		fmt.Println("before middleware")
	}
}

func main(){
	router := gin.Default()

	router.Use(MiddleWare1())
	router.GET("/before", func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"request": request,
		})
	})

	router.Use(MiddleWare2())
	router.GET("/after", func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"request": request,
		})
	})

	router.Run(":8080")
}

/*
curl http://localhost:8080/before
curl http://localhost:8080/after
 */