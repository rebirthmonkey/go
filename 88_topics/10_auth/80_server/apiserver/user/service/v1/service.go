// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/repo"
)

// Service defines functions used to return resource interface.
type Service interface {
	NewUserService() UserService
}

// service is the business logic of the user resource handling.
type service struct {
	repo repo.Repo
}

var _ Service = (*service)(nil)

// NewService returns service instance of the Service interface.
func NewService(repo repo.Repo) Service {
	return &service{repo}
}

// NewUserService returns a user service instance.
func (s *service) NewUserService() UserService {
	return newUserService(s.repo)
}
