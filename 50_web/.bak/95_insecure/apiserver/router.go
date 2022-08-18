package apiserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(g *gin.Engine) {
	installRouterMiddleware(g)
	installController(g)
}

func installRouterMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "GET users")
			})
		}
	}

	return g
}
