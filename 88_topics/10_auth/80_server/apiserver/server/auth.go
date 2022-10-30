// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"time"

	"github.com/rebirthmonkey/go/pkg/auth"

	userRepo "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/repo"
)

type loginInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func newBasicAuth() auth.AuthStrategy {
	return auth.NewBasicStrategy(func(username string, password string) bool {
		user, err := userRepo.Client().UserRepo().Get(username)
		if err != nil {
			return false
		}

		if err := user.Compare(password); err != nil {
			return false
		}

		user.LoginedAt = time.Now()
		_ = userRepo.Client().UserRepo().Update(user)

		return true
	})
}
