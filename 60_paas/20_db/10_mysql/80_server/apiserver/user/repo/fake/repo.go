package fake

import (
	"sync"

	userRepoInterface "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/repo"
)

type repo struct {
	userRepo userRepoInterface.UserRepo
}

var (
	r    userRepoInterface.Repo
	once sync.Once
)

var _ userRepoInterface.Repo = (*repo)(nil)

func Repo() (userRepoInterface.Repo, error) {
	once.Do(func() {
		r = repo{
			userRepo: newUserRepo(),
		}
	})

	return r, nil
}

func (r repo) UserRepo() userRepoInterface.UserRepo {
	return r.userRepo
}

func (r repo) Close() error {
	return nil
}
