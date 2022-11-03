// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	repo3 "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/repo"
	"sync"

	"github.com/rebirthmonkey/go/pkg/mysql"
)

// repo defines the APIServer storage.
type repo struct {
	policyRepo repo3.PolicyRepo
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
			policyRepo: newPolicyRepo(cfg),
		}
	})

	return r, nil
}

// PolicyRepo returns the user store client instance.
func (r repo) PolicyRepo() repo3.PolicyRepo {
	return r.policyRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
