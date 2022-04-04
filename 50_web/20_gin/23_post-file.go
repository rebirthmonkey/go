package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println("filename is:", file.Filename)

		c.SaveUploadedFile(file, "/tmp/"+file.Filename)

		c.String(http.StatusCreated, "upload successful")
	})

	router.Run(":8080")
}
