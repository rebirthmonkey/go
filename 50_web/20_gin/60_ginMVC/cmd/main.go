package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wukong/go/50_web/20_gin/60_ginMVC/controller"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("../template/*")
	router.GET("/hello", hello)

	user := router.Group("/user")
	{
		user.GET("/get/:id/:username", controller.QueryById)
		//user.GET("/query",controller.QueryParam)
		//user.POST("/insert",controller.InsertNewUser)
		//user.GET("/form",controller.RenderForm)
		//user.POST("/form/post",controller.PostForm)
	}

	/*
		file := router.Group("/file")
		{
			// 跳转上传文件页面
			file.GET("/view",controller.RenderView)
			// 根据表单上传
			file.POST("/insert",controller.FormUpload)
			file.POST("/multiUpload",controller.MultiUpload)
			// base64上传
			file.POST("/upload",controller.Base64Upload)
		}
	*/

	router.Run(":8080")
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
	})
}
