// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"sync"

	"github.com/rebirthmonkey/go/pkg/mysql"

	userRepoInterface "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/repo"
)

// repo defines the APIServer storage.
type repo struct {
	userRepo userRepoInterface.UserRepo
}

var (
	r    repo
	once sync.Once
)

var _ userRepoInterface.Repo = (*repo)(nil)

// Repo creates and returns the store client instance.
func Repo(cfg *mysql.CompletedConfig) (userRepoInterface.Repo, error) {
	once.Do(func() {
		r = repo{
			userRepo: newUserRepo(cfg),
		}
	})

	return r, nil
}

// UserRepo returns the user store client instance.
func (r repo) UserRepo() userRepoInterface.UserRepo {
	return r.userRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
