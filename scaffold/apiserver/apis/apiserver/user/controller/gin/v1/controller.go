// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/repo"
	srv "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/service/v1"
)

// Controller creates a user handler interface for user resource.
type Controller interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

// controller creates a user handler used to handle request for user resource.
type controller struct {
	srv srv.Service
}

// NewController creates a user handler.
func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
