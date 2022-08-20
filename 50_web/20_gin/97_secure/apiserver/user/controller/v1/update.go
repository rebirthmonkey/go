// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	model "github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver/user/model/v1"
	"github.com/rebirthmonkey/pkg/gin/util"
	"github.com/rebirthmonkey/pkg/log"
)

func (u *controller) Update(c *gin.Context) {
	log.L(c).Info("[GIN Server] userController: update")

	var m model.User

	if err := c.ShouldBindJSON(&m); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	user, err := u.srv.GetUserService().Get(c.Param("name"))
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	user.Nickname = m.Nickname
	user.Email = m.Email
	user.Phone = m.Phone
	user.Extend = m.Extend

	// Save changed fields.
	if err := u.srv.GetUserService().Update(user); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, user)
}
