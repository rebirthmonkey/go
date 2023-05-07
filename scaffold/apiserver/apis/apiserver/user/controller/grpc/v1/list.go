// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"

	"github.com/rebirthmonkey/go/pkg/log"

	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/controller/grpc/v1/pb"
)

// ListUsers lists the users in the storage.
func (c *controller) ListUsers(ctx context.Context, r *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	log.L(ctx).Info("[GrpcServer] controller: ListUsers")

	users, err := c.srv.NewUserService().List()
	if err != nil {
		return nil, err
	}

	items := make([]*pb.UserInfo, 0)
	for _, user := range users.Items {
		items = append(items, &pb.UserInfo{
			Nickname:  user.Name,
			Password:  user.Password,
			Email:     user.Email,
			Phone:     user.Phone,
			LoginedAt: user.LoginedAt.Format("2006-01-02 15:04:05"),
		})

	}

	return &pb.ListUsersResponse{
		TotalCount: users.TotalCount,
		Items:      items,
	}, nil
}
