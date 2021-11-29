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
		form, _ := c.MultipartForm()
		files := form.File["file[]"]

		for _, file := range files {
			fmt.Println(file.Filename)

			c.SaveUploadedFile(file, "/tmp/xxx-"+file.Filename)
		}

		c.String(http.StatusCreated, "upload successful")
	})

	router.Run(":8080")
}

/*
curl -X POST http://localhost:8080/upload \
  -F "file[]=@./23_post-file.go" \
  -F "file[]=@./24_post-multi-file.go" \
  -H "Content-Type: multipart/form-data"
 */