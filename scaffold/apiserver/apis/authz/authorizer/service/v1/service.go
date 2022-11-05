// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
)

type Service interface {
	NewAuthorizerService() AuthorizerService
}

type service struct {
	repo repo.Repo
}

func NewService(repo repo.Repo) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) NewAuthorizerService() AuthorizerService {
	return newAuthorizerService(s.repo)
}
