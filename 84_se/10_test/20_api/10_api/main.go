package main

// 基于Go Gin的框架开发一个简单的RESTful API, 接口为 HTTP GET http://127.0.0.1:8080/healthz
// 返回JSON格式的数据为 {"message": "hello world"}

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 为GET方法添加一个路由规则，当访问/test时，会执行后面的匿名函数
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// 启动HTTP服务
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("server start failed, err:%v", err)
	}
	return
}
