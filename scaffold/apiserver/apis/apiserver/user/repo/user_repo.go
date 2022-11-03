// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo

import (
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/model/v1"
)

// UserRepo defines the user resources.
type UserRepo interface {
	Create(user *model.User) error
	Delete(username string) error
	Update(user *model.User) error
	Get(username string) (*model.User, error)
	List() (*model.UserList, error)
}
