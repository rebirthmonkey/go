package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/log"

	userCtl "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/controller/v1"
	userRepo "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/repo"
	userRepoFake "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/repo/fake"
)

func InitRouter(g *gin.Engine) {
	installRouterMiddleware(g)
	installController(g)
}

func installRouterMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	//v1 := g.Group("/v1")
	//{
	//	userv1 := v1.Group("/users")
	//	{
	//		userv1.GET("/", func(c *gin.Context) {
	//			c.String(http.StatusOK, "GET users")
	//		})
	//	}
	//}

	v1 := g.Group("/v1")
	{
		log.Info("[GIN Server] registry userHandler")
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
