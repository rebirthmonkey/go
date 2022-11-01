// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"sync"

	"github.com/rebirthmonkey/go/pkg/mysql"

	secretRepoInterface "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/secret/repo"
)

// repo defines the APIServer storage.
type repo struct {
	secretRepo secretRepoInterface.SecretRepo
}

var (
	r    repo
	once sync.Once
)

var _ secretRepoInterface.Repo = (*repo)(nil)

// Repo creates and returns the store client instance.
func Repo(cfg *mysql.CompletedConfig) (secretRepoInterface.Repo, error) {
	once.Do(func() {
		r = repo{
			secretRepo: newSecretRepo(cfg),
		}
	})

	return r, nil
}

// SecretRepo returns the user store client instance.
func (r repo) SecretRepo() secretRepoInterface.SecretRepo {
	return r.secretRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
