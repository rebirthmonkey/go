package mysql

import (
	"sync"

	"github.com/rebirthmonkey/go/pkg/mysql"

	userRepoInterface "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/repo"
)

type repo struct {
	userRepo userRepoInterface.UserRepo
}

var (
	r    repo
	once sync.Once
)

var _ userRepoInterface.Repo = (*repo)(nil)

func Repo(config *mysql.CompletedConfig) (userRepoInterface.Repo, error) {
	once.Do(func() {
		r = repo{
			userRepo: newUserRepo(config),
		}
	})

	return r, nil
}

func (r repo) UserRepo() userRepoInterface.UserRepo {
	return r.userRepo
}

func (r repo) Close() error {
	return r.Close()
}
