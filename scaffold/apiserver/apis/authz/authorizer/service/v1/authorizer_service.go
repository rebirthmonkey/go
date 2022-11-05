// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/ory/ladon"
	"github.com/rebirthmonkey/go/pkg/log"

	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/model/v1"
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
)

type AuthorizerService interface {
	Authorize(*ladon.Request) *model.Response
}

type authorizerService struct {
	warden ladon.Warden
	repo   repo.Repo
}

var _ AuthorizerService = (*authorizerService)(nil)

func newAuthorizerService(repo repo.Repo) AuthorizerService {
	return &authorizerService{
		warden: &ladon.Ladon{
			Manager:     NewLadonManager(repo),
			AuditLogger: &ladon.AuditLoggerInfo{},
		},
		repo: repo,
	}
}

func (a *authorizerService) Authorize(request *ladon.Request) (response *model.Response) {
	log.Info("[Authorizer] start authorization")

	if err := a.warden.IsAllowed(request); err != nil {
		return &model.Response{
			Denied: true,
			Reason: err.Error(),
		}
	}

	return &model.Response{
		Allowed: true,
	}
}
