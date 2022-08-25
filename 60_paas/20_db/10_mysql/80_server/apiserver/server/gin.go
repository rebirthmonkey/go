package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/middleware"

	userCtl "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/controller/gin/v1"
	userRepoMysql "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/repo/mysql"
)

func InitGin(g *gin.Engine) {
	installRouterMiddleware(g)
	installController(g)
}

func installRouterMiddleware(g *gin.Engine) {
	fmt.Println("[GinServer] registry LoggerMiddleware")
	g.Use(middleware.LoggerMiddleware())
}

func installController(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		fmt.Println("[GinServer] registry userHandler")
		userv1 := v1.Group("/users")
		{
			//userRepoClient, _ := userRepoFake.Repo()
			userRepoClient, _ := userRepoMysql.Repo(config.CompletedMysqlConfig)
			userController := userCtl.NewController(userRepoClient)

			userv1.POST("", userController.Create)
			userv1.DELETE(":name", userController.Delete)
			userv1.PUT(":name", userController.Update)
			userv1.GET(":name", userController.Get)
			userv1.GET("", userController.List)
		}
	}
	return g
}
