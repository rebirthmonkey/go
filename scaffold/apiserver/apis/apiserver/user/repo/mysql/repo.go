// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	repo3 "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/repo"
	"sync"

	"github.com/rebirthmonkey/go/pkg/mysql"
)

// repo defines the APIServer storage.
type repo struct {
	userRepo repo3.UserRepo
}

var (
	r    repo
	once sync.Once
)

var _ repo3.Repo = (*repo)(nil)

// Repo creates and returns the store client instance.
func Repo(cfg *mysql.CompletedConfig) (repo3.Repo, error) {
	once.Do(func() {
		r = repo{
			userRepo: newUserRepo(cfg),
		}
	})

	return r, nil
}

// UserRepo returns the user store client instance.
func (r repo) UserRepo() repo3.UserRepo {
	return r.userRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
