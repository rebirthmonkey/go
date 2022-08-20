// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	model "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/model/v1"
	"github.com/rebirthmonkey/pkg/gin/util"
	"github.com/rebirthmonkey/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

func (u *controller) Create(c *gin.Context) {
	log.L(c).Info("[GIN Server] userController: create")

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedBytes)
	user.Status = 1

	if err := u.srv.GetUserService().Create(&user); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, user)
}
