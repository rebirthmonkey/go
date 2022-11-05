// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

func (s *controller) Delete(c *gin.Context) {
	log.L(c).Info("[GINServer] policyController: delete")

	if err := s.srv.NewPolicyService().Delete(c.GetString(UsernameKey), c.Param("name")); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	var msg string = "deleted policy " + c.Param("name")
	util.WriteResponse(c, nil, msg)
}
