// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/pkg/errors"
	"github.com/rebirthmonkey/pkg/gin/util"
)

// Authorize returns whether a request is allow or deny to access a resource and do some action
// under specified condition.
func (a *controller) Authorize(c *gin.Context) {
	var request ladon.Request
	if err := c.ShouldBind(&request); err != nil {
		util.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)
		return
	}

	if request.Context == nil {
		request.Context = ladon.Context{}
	}
	request.Context["username"] = c.GetString("username")

	res := a.srv.NewAuthorizerService().Authorize(&request)

	util.WriteResponse(c, nil, res)
}
