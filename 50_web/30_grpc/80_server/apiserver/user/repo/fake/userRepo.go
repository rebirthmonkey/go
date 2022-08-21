// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fake

import (
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/model/v1"
	model "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/repo"
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/metamodel"
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/reflect"
	"github.com/rebirthmonkey/pkg/errors"
)

type userRepo struct {
	dbEngine []*v1.User
}

var _ userRepoInterface.UserRepo = (*userRepo)(nil)

func newUserRepo(dbEngine []*v1.User) userRepoInterface.UserRepo {
	return &userRepo{dbEngine}
}

func (u *userRepo) Create(user *model.User) error {
	for _, u := range u.dbEngine {
		if u.Name == user.Name {
			return errors.WithCode(88, "record already exist")
		}
	}

	if len(u.dbEngine) > 0 {
		user.ID = u.dbEngine[len(u.dbEngine)-1].ID + 1
	}
	u.dbEngine = append(u.dbEngine, user)

	return nil
}

func (u *userRepo) Delete(username string) error {
	users := u.dbEngine
	u.dbEngine = make([]*v1.User, 0)
	for _, user := range users {
		if user.Name == username {
			continue
		}

		u.dbEngine = append(u.dbEngine, user)
	}

	return nil
}

func (u *userRepo) Update(user *model.User) error {
	for _, u := range u.dbEngine {
		if u.Name == user.Name {
			if _, err := reflect.CopyObj(user, u, nil); err != nil {
				return errors.Wrap(err, "copy user failed")
			}
		}
	}

	return nil
}

func (u *userRepo) Get(username string) (*model.User, error) {
	for _, u := range u.dbEngine {
		if u.Name == username {
			return u, nil
		}
	}

	return nil, errors.WithCode(88, "record not found")

}

func (u *userRepo) List() (*model.UserList, error) {
	users := make([]*v1.User, 0)
	i := 0
	for _, user := range u.dbEngine {
		users = append(users, user)
		i++
	}

	return &v1.UserList{
		ListMeta: metamodel.ListMeta{
			TotalCount: int64(len(u.dbEngine)),
		},
		Items: users,
	}, nil
}
