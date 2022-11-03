// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

func (s *controller) Get(c *gin.Context) {
	log.L(c).Info("[GINServer] policyController: get")

	policy, err := s.srv.NewPolicyService().Get(c.GetString(UsernameKey), c.Param("name"))
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, policy)
}
