package fake

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/metamodel"
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/reflect"

	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/model/v1"
	model "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/model/v1"
	userRepoInterface "github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/user/repo"
)

type userRepo struct {
	dbEngine []*v1.User
}

var _ userRepoInterface.UserRepo = (*userRepo)(nil)

func newUserRepo() userRepoInterface.UserRepo {

	users := make([]*v1.User, 0)
	for i := 1; i <= 10; i++ {
		users = append(users, &v1.User{
			ObjectMeta: metamodel.ObjectMeta{
				Name: fmt.Sprintf("user%d", i),
				ID:   uint64(i),
			},
			Nickname: fmt.Sprintf("user%d", i),
			Password: fmt.Sprintf("User%d@2020", i),
			Email:    fmt.Sprintf("user%d@qq.com", i),
		})
	}

	return &userRepo{
		dbEngine: users,
	}
}

func (u *userRepo) Create(user *model.User) error {
	if len(u.dbEngine) > 0 {
		user.ID = u.dbEngine[len(u.dbEngine)-1].ID + 1
	}
	u.dbEngine = append(u.dbEngine, user)

	return nil
}

func (u *userRepo) Delete(username string) error {
	for _, user := range u.dbEngine {
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
				return err
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

	return nil, errors.New("record not found")

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
