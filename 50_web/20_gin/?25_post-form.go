package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main(){
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}

		filename := header.Filename
		fmt.Println(file, err, filename)

		out, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Println(err)
		}
		c.String(http.StatusCreated, "upload successful")
	})

	router.Run(":8080")
}

/*
curl -X POST http://127.0.0.1:8080/upload -f "upload=@/Users/ruan/workspace/go/50_web/20_gin/abc.txt" -H "Content-Type: multipart/form-data"
*/
