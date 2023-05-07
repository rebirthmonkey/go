// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"

	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/controller/grpc/v1/pb"
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/repo"
	srv "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/service/v1"
)

// Controller creates a GRPC policy interface for policy resource.
type Controller interface {
	ListPolicies(ctx context.Context, r *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error)
}

// controller creates a GRPC policy handler used to handle request for user resource.
type controller struct {
	srv srv.Service

	pb.UnimplementedPolicyServer
}

// NewController creates a GRPC policy handler.
func NewController(repo repo.Repo) *controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
