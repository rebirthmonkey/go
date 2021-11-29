package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main(){
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var user User

		err := c.Bind(&user) // 自动推断content-type bind的是表单还是json
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"username":   user.Username,
			"passwd":     user.Passwd,
			"age":        user.Age,
		})
	})

	router.Run(":8080")
}

/*
curl -X POST http://localhost:8080/login \
	-H "Content-Type:application/json" \
	-d '{"username": "ruan", "passwd": "123", "age": 21}'
 */
