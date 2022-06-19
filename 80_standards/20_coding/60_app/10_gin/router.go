package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(server *GINServer) {
	// middlewares
	server.Engine.Use(LoggerMiddleware())

	// handlers
	if server.healthz {
		server.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"status": "ok",
			})
		})
	}

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
