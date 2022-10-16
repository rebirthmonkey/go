package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/50_web/20_gin/70_swag/controller"
	_ "github.com/rebirthmonkey/go/50_web/20_gin/70_swag/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title Gin Swag 案例
// @version 1.0
// @description This is a Gin Swag server.

// @contact.name wukong
// @contact.url http://www.xxx.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /
func main() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/hello", hello)
	user := router.Group("/user")
	{
		user.GET("/:id/:username", controller.QueryById)
	}
	router.Run(":8080")
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
	})
}
