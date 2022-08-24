package v1

import (
	"context"
	"fmt"
)

func (c *controller) ListUsers(ctx context.Context, r *ListUsersRequest) (*ListUsersResponse, error) {
	fmt.Println("[GrpcServer] controller: ListUsers")

	users, err := c.srv.NewUserService().List()
	if err != nil {
		return nil, err
	}

	items := make([]*UserInfo, 0)
	for _, user := range users.Items {
		items = append(items, &UserInfo{
			Status:      user.Status,
			Nickname:    user.Nickname,
			Password:    user.Password,
			Email:       user.Email,
			Phone:       user.Phone,
			TotalPolicy: user.TotalPolicy,
			LoginedAt:   user.LoginedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &ListUsersResponse{
		TotalCount: users.TotalCount,
		Items:      items,
	}, nil
}
