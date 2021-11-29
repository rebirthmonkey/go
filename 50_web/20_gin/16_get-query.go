package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 所谓query string，即路由用，用?以后连接的key1=value2&key2=value2的形式的参数
	router.GET("/welcome", func(c *gin.Context) {
		// 如果参数不存在，则用缺省值
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8080")
}

/*
curl http://127.0.0.1:8080/welcome
curl http://127.0.0.1:8080/welcome\?firstname\=中国
curl http://127.0.0.1:8080/welcome\?firstname\=中国\&lastname\=天朝
curl http://127.0.0.1:8080/welcome\?firstname\=\&lastname\=天朝
curl http://127.0.0.1:8080/welcome\?firstname\=%E4%B8%AD%E5%9B%BD
*/