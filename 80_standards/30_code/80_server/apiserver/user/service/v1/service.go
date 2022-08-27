package v1

import (
	"github.com/rebirthmonkey/go/80_standards/30_code/80_server/apiserver/user/repo"
)

type Service interface {
	NewUserService() UserService
}

type service struct {
	repo repo.Repo
}

var _ Service = (*service)(nil)

func NewService(repo repo.Repo) Service {
	return &service{repo}
}

func (s *service) NewUserService() UserService {
	return newUserService(s.repo)
}
