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
		log.Info("[GinServer] registry authzHandler")
		authzv1 := v1.Group("/authz")
		{
			authzRepoClient, _ := authzRepoRest.Repo(config.CompletedAuthzConfig)
			authzRepo.SetClient(authzRepoClient)
			authzController := authzCtl.NewController(authzRepoClient)
			authzv1.POST("", authzController.Authorize)
		}
	}

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)

	v2 := g.Group("/v2")
	{
		log.Info("[GinServer] registry authzHandler")
		authzv2 := v2.Group("/authz", jwtStrategy.AuthFunc())
		{
			authzRepoClient, _ := authzRepoRest.Repo(config.CompletedAuthzConfig)
			authzRepo.SetClient(authzRepoClient)
			authzController := authzCtl.NewController(authzRepoClient)
			authzv2.POST("", authzController.Authorize)
		}
	}

	return g
}
