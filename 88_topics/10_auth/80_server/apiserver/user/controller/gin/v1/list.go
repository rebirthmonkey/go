// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

// List lists the users in the storage.
// Only administrator can call this function.
func (u *controller) List(c *gin.Context) {
	log.L(c).Info("[GinServer] userController: list")

	users, err := u.srv.NewUserService().List()
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, users)
}
