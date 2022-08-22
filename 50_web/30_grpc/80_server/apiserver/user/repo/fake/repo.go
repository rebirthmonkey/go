package fake

import (
	userRepoInterface "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/repo"
)

type repo struct {
	userRepo userRepoInterface.UserRepo
}

var _ userRepoInterface.Repo = (*repo)(nil)

func NewRepo() (userRepoInterface.Repo, error) {
	r := &repo{
		userRepo: newUserRepo(),
	}
	return r, nil
}

func (r *repo) UserRepo() userRepoInterface.UserRepo {
	return r.userRepo
}

func (r *repo) Close() error {
	return nil
}
