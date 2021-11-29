package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
		contentType := c.Request.Header.Get("Content-Type")

		switch contentType {
		case "application/json":
			_ = c.BindJSON(&user)
		case "application/x-www-form-urlencoded":
			_ = c.BindWith(&user, binding.Form)
		}

		c.JSON(http.StatusOK, gin.H{
			"user":   user.Username,
			"passwd": user.Passwd,
			"age":    user.Age,
		})
	})

	router.Run(":8080")
}

/*
curl -X POST http://localhost:8080/login \
	-H "Content-Type:application/json" \
	-d '{"username": "ruan", "passwd": "123", "age": 21}'
 */
