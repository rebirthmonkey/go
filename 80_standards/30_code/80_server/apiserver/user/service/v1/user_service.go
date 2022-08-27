package v1

import (
	"time"

	"github.com/rebirthmonkey/go/pkg/metamodel"
	"golang.org/x/crypto/bcrypt"

	model "github.com/rebirthmonkey/go/80_standards/30_code/80_server/apiserver/user/model/v1"
	"github.com/rebirthmonkey/go/80_standards/30_code/80_server/apiserver/user/repo"
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
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedBytes)
	user.Status = 1
	user.LoginedAt = time.Now()

	return u.repo.UserRepo().Create(user)
}

func (u *userService) Delete(username string) error {
	return u.repo.UserRepo().Delete(username)
}

func (u *userService) Update(user *model.User) error {
	updateUser, err := u.Get(user.Name)
	if err != nil {
		return err
	}

	updateUser.Nickname = user.Nickname
	updateUser.Email = user.Email
	updateUser.Phone = user.Phone
	updateUser.Extend = user.Extend

	return u.repo.UserRepo().Update(updateUser)
}

func (u *userService) Get(username string) (*model.User, error) {
	return u.repo.UserRepo().Get(username)
}

func (u *userService) List() (*model.UserList, error) {
	users, err := u.repo.UserRepo().List()
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
