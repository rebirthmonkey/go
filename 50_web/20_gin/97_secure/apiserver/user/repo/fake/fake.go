// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fake

import (
	"github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver/user/repo"
)

type repo struct {
	dbEngine []*v1.User
}

var _ userRepoInterface.Repo = (*repo)(nil)

func NewRepo(opts interface{}) (userRepoInterface.Repo, error) {
	return &repo{}, nil
}

func (r *repo) GetUserRepo() userRepoInterface.UserRepo {
	return newUserRepo(r.dbEngine)
}

func (r *repo) Close() error {
	r.dbEngine = nil

	return nil
}
