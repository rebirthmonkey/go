package repo

import (
	model "github.com/rebirthmonkey/go/80_standards/20_error/80_server/apiserver/user/model/v1"
)

// UserRepo defines the user resources.
type UserRepo interface {
	Create(user *model.User) error
	Delete(username string) error
	Update(user *model.User) error
	Get(username string) (*model.User, error)
	List() (*model.UserList, error)
}