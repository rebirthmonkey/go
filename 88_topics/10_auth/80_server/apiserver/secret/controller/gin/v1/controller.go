// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/secret/repo"
	srv "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/secret/service/v1"
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

// controller creates a secret handler used to handle request for secret resource.
type controller struct {
	srv srv.Service
}

// NewController creates a secret handler.
func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
