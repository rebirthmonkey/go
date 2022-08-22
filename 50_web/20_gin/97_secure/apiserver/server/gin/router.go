package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/middleware"

	userCtl "github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver/user/controller/v1"
	userRepoFake "github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver/user/repo/fake"
)

func InitRouter(g *gin.Engine) {
	installRouterMiddleware(g)
	installController(g)
}

func installRouterMiddleware(g *gin.Engine) {
	fmt.Println("[GINServer] registry LoggerMiddleware")
	g.Use(middleware.LoggerMiddleware())
}

func installController(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		fmt.Println("[GINServer] registry userHandler")
		userv1 := v1.Group("/users")
		{
			userRepoClient, _ := userRepoFake.NewRepo()
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
