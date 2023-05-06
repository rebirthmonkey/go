// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/auth"
	"github.com/rebirthmonkey/go/pkg/log"
	authzCtl "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/controller/v1"
	authzRepo "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
	authzRepoRest "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo/rest"
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
		//log.Info("[GinServer] registry policyHandler")
		//policyv1 := v1.Group("/policies")
		////v1.Group("/policies")
		//{
		//	policyRepoClient, _ := policyRepoMysql.Repo(config.CompletedMysqlConfig)
		//
		//	policyController := policyCtl.NewController(policyRepoClient)
		//	//policyCtl.NewController(policyRepoClient)
		//
		//	//policyv1.POST("", policyController.Create)
		//	//policyv1.DELETE(":name", policyController.Delete)
		//	//policyv1.PUT(":name", policyController.Update)
		//	//policyv1.GET(":name", policyController.Get)
		//	policyv1.GET("", policyController.List)
		//}

		log.Info("[GinServer] registry authzHandler")
		authzv1 := v1.Group("/authz")
		{
			authzRepoClient, _ := authzRepoRest.Repo(config.CompletedGinConfig)
			authzRepo.SetClient(authzRepoClient)
			authzController := authzCtl.NewController(authzRepoClient)
			authzv1.POST("", authzController.Authorize)
		}
	}

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)

	//v2 := g.Group("/v2")
	v2 := g.Group("/v2")
	{
		//log.Info("[GinServer] registry policyHandler")
		//policyv2 := v2.Group("/policies")
		////v2.Group("/policies")
		//{
		//	policyRepoClient, _ := policyRepoMysql.Repo(config.CompletedMysqlConfig)
		//
		//	policyController := policyCtl.NewController(policyRepoClient)
		//	//policyCtl.NewController(policyRepoClient)
		//
		//	//policyv2.POST("", policyController.Create)
		//	//policyv2.DELETE(":name", policyController.Delete)
		//	//policyv2.PUT(":name", policyController.Update)
		//	//policyv2.GET(":name", policyController.Get)
		//	policyv2.GET("", policyController.List)
		//}

		log.Info("[GinServer] registry authzHandler")
		authzv2 := v2.Group("/authz", jwtStrategy.AuthFunc())
		{
			authzRepoClient, _ := authzRepoRest.Repo(config.CompletedGinConfig)
			authzRepo.SetClient(authzRepoClient)
			authzController := authzCtl.NewController(authzRepoClient)
			authzv2.POST("", authzController.Authorize)
		}
	}

	return g
}
