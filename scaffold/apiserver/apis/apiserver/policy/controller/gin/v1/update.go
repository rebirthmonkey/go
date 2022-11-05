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
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/model/v1"
)

func (s *controller) Update(c *gin.Context) {
	log.L(c).Info("[GINServer] policyController: update")

	var r model.Policy

	if err := c.ShouldBindJSON(&r); err != nil {
		util.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)
		return
	}

	username := c.GetString(UsernameKey)
	policyName := c.Param("name")

	policy, err := s.srv.NewPolicyService().Get(username, policyName)
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	policy.AuthzPolicy = r.AuthzPolicy
	policy.Extend = r.Extend

	if err := s.srv.NewPolicyService().Update(policy); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, policy)
}
