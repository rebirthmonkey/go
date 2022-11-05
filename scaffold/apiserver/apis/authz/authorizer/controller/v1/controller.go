// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
	srv "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/service/v1"
)

type Controller interface {
	Authorize(c *gin.Context)
}

type controller struct {
	srv srv.Service
}

func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
