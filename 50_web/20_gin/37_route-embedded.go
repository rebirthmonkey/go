package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)

func main() {
	router := gin.Default()

	//设置路由前缀 调用Group方法
	userGroup := router.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user/index",
			})
		})

		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user/login",
				})
			})

		shopGroup := userGroup.Group("/shop")
		{
			shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/user/shop/index",
				})
			})
		}
	}
	router.Run()
}

/*
curl http://localhost:8080/user/index
curl http://localhost:8080/user/login
curl http://localhost:8080/user/shop/index
 */