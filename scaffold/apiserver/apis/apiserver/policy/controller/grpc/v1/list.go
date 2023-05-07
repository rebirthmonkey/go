// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"

	"github.com/rebirthmonkey/go/pkg/log"

	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/controller/grpc/v1/pb"
)

// ListPolicies lists the policies in the storage.
func (c *controller) ListPolicies(ctx context.Context, r *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	log.L(ctx).Info("[GrpcServer] controller: ListPolicies")

	policies, err := c.srv.NewPolicyService().List()
	if err != nil {
		return nil, err
	}

	items := make([]*pb.PolicyInfo, 0)
	for _, policy := range policies.Items {
		items = append(items, &pb.PolicyInfo{
			Username:     policy.Username,
			AuthzPolicy:  policy.AuthzPolicy.String(),
			PolicyShadow: policy.PolicyShadow,
		})

	}

	return &pb.ListPoliciesResponse{
		TotalCount: policies.TotalCount,
		Items:      items,
	}, nil
}
