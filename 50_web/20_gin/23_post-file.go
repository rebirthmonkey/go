package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginEngine := gin.Default()
	ginEngine.MaxMultipartMemory = 8 << 20

	ginEngine.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println("filename is:", file.Filename)

		c.SaveUploadedFile(file, "/tmp/"+file.Filename)

		c.String(http.StatusCreated, "upload successful")
	})

	ginEngine.Run(":8080")
}
