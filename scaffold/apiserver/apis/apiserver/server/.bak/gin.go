// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/auth"
	"github.com/rebirthmonkey/go/pkg/gin/middleware"
	"github.com/rebirthmonkey/go/pkg/log"
	policyCtl "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/controller/gin/v1"
	policyRepo "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/repo"
	policyRepoMysql "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/repo/mysql"
	secretCtl "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/secret/controller/gin/v1"
	secretRepo "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/secret/repo"
	secretRepoMysql "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/secret/repo/mysql"
	userCtl "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/controller/gin/v1"
	userRepo "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/repo"
	userRepoMysql "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/repo/mysql"
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

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)

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
			userRepo.SetClient(userRepoClient)

			userController := userCtl.NewController(userRepoClient)

			basicStrategy := newBasicAuth()
			userv1.Use(basicStrategy.AuthFunc())

			userv1.POST("", userController.Create)
			userv1.DELETE(":name", userController.Delete)
			userv1.PUT(":name", userController.Update)
			userv1.GET(":name", userController.Get)
			userv1.GET("", userController.List)
		}

		log.Info("[GINServer] registry secretHandler")
		//secretv1 := v1.Group("/secrets")
		//secretv1.Use(jwtStrategy.AuthFunc())
		secretv1 := v1.Group("/secrets", jwtStrategy.AuthFunc())
		{
			secretRepoClient, _ := secretRepoMysql.Repo(config.CompletedMysqlConfig)
			secretRepo.SetClient(secretRepoClient)

			secretController := secretCtl.NewController(secretRepoClient)

			secretv1.POST("", secretController.Create)
			secretv1.DELETE(":name", secretController.Delete)
			secretv1.PUT(":name", secretController.Update)
			secretv1.GET(":name", secretController.Get)
			secretv1.GET("", secretController.List)
		}

		log.Info("[GINServer] registry policyHandler")
		policyv1 := v1.Group("/policies")
		{
			policyRepoClient, _ := policyRepoMysql.Repo(config.CompletedMysqlConfig)
			policyRepo.SetClient(policyRepoClient)

			policyController := policyCtl.NewController(policyRepoClient)

			policyv1.POST("", policyController.Create)
			policyv1.DELETE(":name", policyController.Delete)
			policyv1.PUT(":name", policyController.Update)
			policyv1.GET(":name", policyController.Get)
			policyv1.GET("", policyController.List)
		}
	}
	return g
}
