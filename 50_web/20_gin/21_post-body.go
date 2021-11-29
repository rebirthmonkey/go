package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick": nick,
		})
	})

	router.Run(":8080")
}

/*
curl -X POST http://127.0.0.1:8080/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=ruan"
 */