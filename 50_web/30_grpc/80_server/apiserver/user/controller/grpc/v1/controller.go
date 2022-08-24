package v1

import (
	"context"

	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/repo"
	srv "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/service/v1"
)

type Controller interface {
	ListUsers(ctx context.Context, r *ListUsersRequest) (*ListUsersResponse, error)
}

type controller struct {
	srv srv.Service
	UnimplementedUserServer
}

func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
func (c *controller) mustEmbedUnimplementedUserServer() {
	//TODO implement me
	panic("implement me")
}
