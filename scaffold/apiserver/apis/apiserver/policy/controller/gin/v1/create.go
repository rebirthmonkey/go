// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	ginUtil "github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/model/v1"
)

// Create add new policy to the storage.
func (s *controller) Create(c *gin.Context) {
	log.L(c).Info("[GINServer] policyController: create")

	var policy model.Policy

	if err := c.ShouldBindJSON(&policy); err != nil {
		ginUtil.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)
		return
	}

	policy.Username = c.GetString(UsernameKey)

	if err := s.srv.NewPolicyService().Create(&policy); err != nil {
		ginUtil.WriteResponse(c, err, nil)
		return
	}

	ginUtil.WriteResponse(c, nil, policy)
}
