// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	repo3 "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/secret/repo"
	"sync"

	"github.com/rebirthmonkey/go/pkg/mysql"
)

// repo defines the APIServer storage.
type repo struct {
	secretRepo repo3.SecretRepo
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
			secretRepo: newSecretRepo(cfg),
		}
	})

	return r, nil
}

// SecretRepo returns the user store client instance.
func (r repo) SecretRepo() repo3.SecretRepo {
	return r.secretRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
