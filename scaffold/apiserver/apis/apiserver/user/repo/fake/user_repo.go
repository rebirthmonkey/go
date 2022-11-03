// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fake

import (
	"fmt"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/metamodel"
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/repo"
)

// userRepo stores the user's info.
type userRepo struct {
	dbEngine []*model.User
}

var _ userRepoInterface.UserRepo = (*userRepo)(nil)

// newUserRepo creates and returns a user storage.
func newUserRepo() userRepoInterface.UserRepo {

	users := make([]*model.User, 0)
	for i := 1; i <= 10; i++ {
		users = append(users, &model.User{
			ObjectMeta: metamodel.ObjectMeta{
				Name: fmt.Sprintf("user%d", i),
				ID:   uint64(i),
			},
			Nickname: fmt.Sprintf("user%d", i),
			Password: fmt.Sprintf("User%d@2022", i),
			Email:    fmt.Sprintf("user%d@qq.com", i),
		})
	}

	return &userRepo{
		dbEngine: users,
	}
}

// Create creates a new user account.
func (u *userRepo) Create(user *model.User) error {
	for _, u := range u.dbEngine {
		if u.Name == user.Name {
			return errors.WithCode(errcode.ErrRecordAlreadyExist, "record already exist")
		}
	}

	if len(u.dbEngine) > 0 {
		user.ID = u.dbEngine[len(u.dbEngine)-1].ID + 1
	}
	u.dbEngine = append(u.dbEngine, user)

	return nil
}

// Delete deletes the user by the user identifier.
func (u *userRepo) Delete(username string) error {
	newUsers := make([]*model.User, 0)

	for i := 0; i < len(u.dbEngine); i++ {
		if u.dbEngine[i].Name == username {
			newUsers = append(u.dbEngine[:i], u.dbEngine[i+1:]...)
			break
		}
	}

	if len(newUsers) == 0 {
		return errors.WithCode(errcode.ErrRecordNotFound, "record not found")
	}

	u.dbEngine = newUsers
	return nil
}

// Update updates a user account information.
func (u *userRepo) Update(user *model.User) error {
	if err := u.Delete(user.Name); err != nil {
		return err
	}

	return u.Create(user)
}

// Get returns a user's info by the user identifier.
func (u *userRepo) Get(username string) (*model.User, error) {
	for _, u := range u.dbEngine {
		if u.Name == username {
			return u, nil
		}
	}

	return nil, errors.WithCode(errcode.ErrRecordNotFound, "record not found")
}

// List returns all the related users.
func (u *userRepo) List() (*model.UserList, error) {
	users := make([]*model.User, 0)
	i := 0
	for _, user := range u.dbEngine {
		users = append(users, user)
		i++
	}

	return &model.UserList{
		ListMeta: metamodel.ListMeta{
			TotalCount: int64(len(u.dbEngine)),
		},
		Items: users,
	}, nil
}
