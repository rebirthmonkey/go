package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main(){
	router := gin.Default()

	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path" + c.Request.URL.Path)
	})

	router.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})

	router.Run(":8080")
}

