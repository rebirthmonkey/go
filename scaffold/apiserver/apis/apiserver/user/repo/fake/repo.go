// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fake

import (
	repo3 "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/repo"
	"sync"
)

// repo defines the APIServer storage.
type repo struct {
	userRepo repo3.UserRepo
}

var (
	r    repo3.Repo
	once sync.Once
)

var _ repo3.Repo = (*repo)(nil)

// Repo creates and returns the store client instance.
func Repo() (repo3.Repo, error) {
	once.Do(func() {
		r = repo{
			userRepo: newUserRepo(),
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
	return nil
}
