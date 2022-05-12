package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User2 struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main() {
	ginEngine := gin.Default()

	ginEngine.POST("/login", func(c *gin.Context) {
		var user User2

		err := c.Bind(&user) // 自动推断content-type bind的是表单还是json
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"passwd":   user.Passwd,
			"age":      user.Age,
		})
	})

	ginEngine.Run(":8080")
}
