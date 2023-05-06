// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/auth"
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
	//g.Use(middleware.LoggerMiddleware())
}

// installController installs Gin handlers
func installController(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		log.Info("[GinServer] registry userHandler")
		userv1 := v1.Group("/users")
		{
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

		secretv1 := v1.Group("/secrets")
		{
			secretRepoClient, err := secretRepoMysql.Repo(config.CompletedMysqlConfig)
			if err != nil {
				log.Fatalf("failed to create Mysql repo: %s", err.Error())
			}

			secretController := secretCtl.NewController(secretRepoClient)

			secretv1.POST("", secretController.Create)
			secretv1.DELETE(":name", secretController.Delete)
			secretv1.PUT(":name", secretController.Update)
			secretv1.GET(":name", secretController.Get)
			secretv1.GET("", secretController.List)
		}

		policyv1 := v1.Group("/policies")
		{
			policyRepoClient, err := policyRepoMysql.Repo(config.CompletedMysqlConfig)
			if err != nil {
				log.Fatalf("failed to create Mysql repo: %s", err.Error())
			}

			policyController := policyCtl.NewController(policyRepoClient)

			policyv1.POST("", policyController.Create)
			policyv1.DELETE(":name", policyController.Delete)
			policyv1.PUT(":name", policyController.Update)
			policyv1.GET(":name", policyController.Get)
			policyv1.GET("", policyController.List)
		}
	}

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)

	v2 := g.Group("/v2")
	{
		log.Info("[GinServer] registry userHandler with basic auth")
		basicStrategy := newBasicAuth()
		userv2 := v2.Group("/users", basicStrategy.AuthFunc())
		{
			userRepoClient, _ := userRepoMysql.Repo(config.CompletedMysqlConfig)
			userRepo.SetClient(userRepoClient)

			userController := userCtl.NewController(userRepoClient)

			userv2.POST("", userController.Create)
			userv2.DELETE(":name", userController.Delete)
			userv2.PUT(":name", userController.Update)
			userv2.GET(":name", userController.Get)
			userv2.GET("", userController.List)
		}

		log.Info("[GINServer] registry secretHandler with jwt auth")
		secretv2 := v2.Group("/secrets", jwtStrategy.AuthFunc())
		{
			secretRepoClient, _ := secretRepoMysql.Repo(config.CompletedMysqlConfig)
			secretRepo.SetClient(secretRepoClient)

			secretController := secretCtl.NewController(secretRepoClient)

			secretv2.POST("", secretController.Create)
			secretv2.DELETE(":name", secretController.Delete)
			secretv2.PUT(":name", secretController.Update)
			secretv2.GET(":name", secretController.Get)
			secretv2.GET("", secretController.List)
		}

		log.Info("[GINServer] registry policyHandler with jwt auth")
		policyv2 := v2.Group("/policies", jwtStrategy.AuthFunc())
		{
			policyRepoClient, _ := policyRepoMysql.Repo(config.CompletedMysqlConfig)
			policyRepo.SetClient(policyRepoClient)

			policyController := policyCtl.NewController(policyRepoClient)

			policyv2.POST("", policyController.Create)
			policyv2.DELETE(":name", policyController.Delete)
			policyv2.PUT(":name", policyController.Update)
			policyv2.GET(":name", policyController.Get)
			policyv2.GET("", policyController.List)
		}
	}
	return g
}
