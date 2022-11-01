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

	model "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/secret/model/v1"
)

func (s *controller) Update(c *gin.Context) {
	log.L(c).Info("[GINServer] secretController: update")

	var r model.Secret
	if err := c.ShouldBindJSON(&r); err != nil {
		util.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)
		return
	}

	username := c.GetString(UsernameKey)
	secretName := c.Param("name")

	secret, err := s.srv.NewSecretService().Get(username, secretName)
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	secret.Expires = r.Expires
	secret.Description = r.Description
	secret.Extend = r.Extend

	if err2 := s.srv.NewSecretService().Update(secret); err2 != nil {
		util.WriteResponse(c, err2, nil)

		return
	}

	util.WriteResponse(c, nil, secret)
}
