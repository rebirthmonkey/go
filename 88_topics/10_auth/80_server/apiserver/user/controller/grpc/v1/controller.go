// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"

	"github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/repo"
	srv "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/service/v1"
)

// Controller creates a GRPC user interface for user resource.
type Controller interface {
	ListUsers(ctx context.Context, r *ListUsersRequest) (*ListUsersResponse, error)
}

// controller creates a GRPC user handler used to handle request for user resource.
type controller struct {
	srv srv.Service
	UnimplementedUserServer
}

// NewController creates a GRPC user handler.
func NewController(repo repo.Repo) *controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
