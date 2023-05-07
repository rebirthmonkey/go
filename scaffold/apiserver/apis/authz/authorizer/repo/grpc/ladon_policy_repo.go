// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"encoding/json"
	"github.com/avast/retry-go"
	"sync"

	"github.com/ory/ladon"
	"github.com/rebirthmonkey/go/pkg/log"
	"google.golang.org/grpc"

	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/controller/grpc/v1/pb"
	authzRepo "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/server/authz"
)

type ladonPolicyRepo struct {
	policyURL string
}

var (
	landonPolicyRepoInstance *ladonPolicyRepo
	onceCache                sync.Once
)

var _ authzRepo.LadonPolicyRepo = (*ladonPolicyRepo)(nil)

func newLadonPolicyRepo(cfg *authz.CompletedConfig) authzRepo.LadonPolicyRepo {
	onceCache.Do(func() {
		landonPolicyRepoInstance = &ladonPolicyRepo{
			policyURL: cfg.PolicyServer,
		}
	})

	return landonPolicyRepoInstance
}

func (p *ladonPolicyRepo) List() ([]*ladon.DefaultPolicy, error) {
	conn, err := grpc.Dial(p.policyURL, grpc.WithInsecure())
	if err != nil {
		log.Error("[AuthzPolicy/repo/grpc] List: GRPC Response error")
		return nil, err
	}
	defer conn.Close()

	client := pb.NewPolicyClient(conn)
	var resp *pb.ListPoliciesResponse

	err = retry.Do(
		func() error {
			var listErr error
			resp, listErr = client.ListPolicies(context.TODO(), &pb.ListPoliciesRequest{})
			if listErr != nil {
				return listErr
			}

			return nil
		}, retry.Attempts(3),
	)

	var ladonPolicyList []*ladon.DefaultPolicy

	for _, v := range resp.Items {
		policy := ladon.DefaultPolicy{}

		if err := json.Unmarshal([]byte(v.PolicyShadow), &policy); err != nil {
			log.Errorf("failed to load policy with error: %s", err.Error())

			continue
		}

		ladonPolicyList = append(ladonPolicyList, &policy)
	}

	return ladonPolicyList, nil
}
