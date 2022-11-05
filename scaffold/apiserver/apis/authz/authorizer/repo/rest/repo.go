// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"github.com/rebirthmonkey/go/pkg/gin"

	authzRepo "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
	"sync"
)

type repo struct {
	ladonPolicyRepo authzRepo.LadonPolicyRepo
}

var (
	r    repo
	once sync.Once
)

func Repo(cfg *gin.CompletedConfig) (authzRepo.Repo, error) {
	once.Do(func() {
		r = repo{
			ladonPolicyRepo: newLadonPolicyRepo(cfg),
		}
	})

	return r, nil
}

// LadonPolicyRepo returns the ladon policy repo client instance.
func (r repo) LadonPolicyRepo() authzRepo.LadonPolicyRepo {
	return r.ladonPolicyRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
