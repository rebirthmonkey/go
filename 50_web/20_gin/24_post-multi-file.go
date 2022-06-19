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
		form, _ := c.MultipartForm()
		files := form.File["file[]"]

		for _, file := range files {
			fmt.Println(file.Filename)

			c.SaveUploadedFile(file, "/tmp/xxx-"+file.Filename)
		}

		c.String(http.StatusCreated, "upload successful")
	})

	ginEngine.Run(":8080")
}
