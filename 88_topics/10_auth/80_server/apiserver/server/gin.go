// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/middleware"
	"github.com/rebirthmonkey/go/pkg/log"

	userCtl "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/controller/gin/v1"
	userRepoMysql "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/repo/mysql"
)

// InitGin initializes the Gin server
func InitGin(g *gin.Engine) {
	installRouterMiddleware(g)
	installController(g)
}

// installRouterMiddleware installs Gin server middlewares
func installRouterMiddleware(g *gin.Engine) {
	log.Info("[GinServer] registry LoggerMiddleware")
	g.Use(middleware.LoggerMiddleware())
}

// installController installs Gin handlers
func installController(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		log.Info("[GinServer] registry userHandler")
		userv1 := v1.Group("/users")
		{
			//userRepoClient, err := userRepoFake.Repo()
			//if err != nil {
			//	log.Fatalf("failed to create fake repo: %s", err.Error())
			//}

			userRepoClient, err := userRepoMysql.Repo(config.CompletedMysqlConfig)
			if err != nil {
				log.Fatalf("failed to create Mysql repo: %s", err.Error())
			}

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
