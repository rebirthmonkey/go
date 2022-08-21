package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/log"

	userCtl "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/controller/v1"
	userRepo "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/repo"
	userRepoFake "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/repo/fake"
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/gin/middleware"
)

func Init(g *gin.Engine) {
	installRouterMiddleware(g)
	installController(g)
}

func installRouterMiddleware(g *gin.Engine) {
	log.Info("[GinServer] registry LoggerMiddleware")
	g.Use(middleware.LoggerMiddleware())
}

func installController(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		log.Info("[GinServer] registry userHandler")
		//mysqlOptions := mysql.NewOptions(optionFile)
		userv1 := v1.Group("/users")
		{
			userRepoFakeInstance, _ := userRepoFake.NewRepo(nil)

			userRepo.SetRepo(userRepoFakeInstance)
			userController := userCtl.NewController(userRepoFakeInstance)

			userv1.POST("", userController.Create)
			userv1.DELETE(":name", userController.Delete)
			userv1.PUT(":name", userController.Update)
			userv1.GET(":name", userController.Get)
			userv1.GET("", userController.List)
		}
	}
	return g
}
