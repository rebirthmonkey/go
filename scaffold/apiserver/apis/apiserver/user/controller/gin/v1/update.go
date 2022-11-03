// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/model/v1"
)

// Update updates a user's info by the user identifier.
func (u *controller) Update(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: update")

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.L(c).Errorf("ErrBind: %s\n", err)
		util.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)

		return
	}

	user.Name = c.Param("name")

	if err := u.srv.NewUserService().Update(&user); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, user)
}
