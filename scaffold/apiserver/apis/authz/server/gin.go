// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/middleware"
	"github.com/rebirthmonkey/go/pkg/log"

	authorizerCtl "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/controller/v1"
	authorizerRepo "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
	authorizerRepoRest "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo/rest"
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
		log.Info("[GinServer] registry authorizer")
		authzv1 := v1.Group("/authz")
		{
			authzRepoClient, err := authorizerRepoRest.Repo(config.CompletedApiserverConfig)
			if err != nil {
				log.Fatalf("failed to create REST repo: %s", err.Error())
			}
			authorizerRepo.SetClient(authzRepoClient)

			authorizerController := authorizerCtl.NewController(authzRepoClient)

			authzv1.POST("", authorizerController.Authorize)
		}
	}
	return g
}
