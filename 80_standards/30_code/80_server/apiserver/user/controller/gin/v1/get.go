// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

// Get gets a user's info by the user identifier.
func (u *controller) Get(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: get")

	user, err := u.srv.NewUserService().Get(c.Param("name"))
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, user)
}
