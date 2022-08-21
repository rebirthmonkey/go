// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/gin/util"
	"github.com/rebirthmonkey/pkg/log"
)

func (u *controller) Delete(c *gin.Context) {
	log.L(c).Info("[GIN Server] userController: delete")

	if err := u.srv.NewUserService().Delete(c.Param("name")); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, nil)
}
