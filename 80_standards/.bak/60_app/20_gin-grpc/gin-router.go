package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/log"
	"net/http"
)

func initRouter(server *GINServer) {
	// middlewares
	server.Engine.Use(LoggerMiddleware())
	log.Info("[GIN Server] registry LoggerMiddleware")


	// handlers
	log.Info("[GIN Server] registry healthz")
	if server.healthz {
		server.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"status": "ok",
			})
		})
	}

	log.Info("[GIN Server] registry productHandler")
	v1 := server.Engine.Group("/v1")
	{
		productv1 := v1.Group("/products")
		productHandler := newProductHandler()
		{
			productv1.POST("", productHandler.Create)
			productv1.GET(":name", productHandler.Get)
		}
	}
}
