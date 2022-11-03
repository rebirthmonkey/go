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
	"github.com/rebirthmonkey/go/pkg/util"
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/secret/model/v1"
)

// Create add new secret to the storage.
func (s *controller) Create(c *gin.Context) {
	log.L(c).Info("[GINServer] secretController: create")

	var secret model.Secret

	if err := c.ShouldBindJSON(&secret); err != nil {
		ginUtil.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)
		return
	}

	//secret.Username = c.GetString(UsernameKey)
	secret.SecretID = util.NewSecretID()
	secret.SecretKey = util.NewSecretKey()

	if err := s.srv.NewSecretService().Create(&secret); err != nil {
		ginUtil.WriteResponse(c, err, nil)
		return
	}

	ginUtil.WriteResponse(c, nil, secret)
}
