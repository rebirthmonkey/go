// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver/user/repo"
)

type Service interface {
	GetUserService() UserService
}

type service struct {
	repo repo.Repo
}

var _ Service = (*service)(nil)

func NewService(repo repo.Repo) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUserService() UserService {
	return newUserService(s.repo)
}
