// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/repo"
	srv "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/service/v1"
)

const UsernameKey = "username"

// Controller creates a secret handler interface for secret resource.
type Controller interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

// controller creates a policy handler used to handle request for policy resource.
type controller struct {
	srv srv.Service
}

// NewController creates a secret handler.
func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
