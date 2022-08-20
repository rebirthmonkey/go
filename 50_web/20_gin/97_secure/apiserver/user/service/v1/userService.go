// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"time"

	model "github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver/user/model/v1"
	"github.com/rebirthmonkey/go/50_web/20_gin/97_secure/apiserver/user/repo"
	"github.com/rebirthmonkey/go/50_web/20_gin/97_secure/pkg/metamodel"
)

type UserService interface {
	Create(user *model.User) error
	Delete(username string) error
	Update(user *model.User) error
	Get(username string) (*model.User, error)
	List() (*model.UserList, error)
}

type userService struct {
	repo repo.Repo
}

var _ UserService = (*userService)(nil)

func newUserService(repo repo.Repo) UserService {
	return &userService{repo}
}

func (u *userService) Create(user *model.User) error {
	user.LoginedAt = time.Now()
	return u.repo.GetUserRepo().Create(user)
}

func (u *userService) Delete(username string) error {
	return u.repo.GetUserRepo().Delete(username)
}

func (u *userService) Update(user *model.User) error {
	return u.repo.GetUserRepo().Update(user)
}

func (u *userService) Get(username string) (*model.User, error) {
	return u.repo.GetUserRepo().Get(username)
}

func (u *userService) List() (*model.UserList, error) {
	users, err := u.repo.GetUserRepo().List()
	if err != nil {
		return nil, err
	}

	infos := make([]*model.User, 0)
	for _, user := range users.Items {
		infos = append(infos, &model.User{
			ObjectMeta: metamodel.ObjectMeta{
				ID:        user.ID,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
		})
	}

	return &model.UserList{ListMeta: users.ListMeta, Items: infos}, nil
}
