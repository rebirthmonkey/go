// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/model/v1"
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/repo"
)

type PolicyService interface {
	Create(policy *model.Policy) error
	Delete(username, policyName string) error
	Update(policy *model.Policy) error
	Get(username, policyName string) (*model.Policy, error)
	List(username string) (*model.PolicyList, error)
}

type policyService struct {
	repo repo.Repo
}

var _ PolicyService = (*policyService)(nil)

// newPolicyService creates and returns the policy service instance.
func newPolicyService(repo repo.Repo) PolicyService {
	return &policyService{repo: repo}
}

func (p *policyService) Create(policy *model.Policy) error {
	return p.repo.PolicyRepo().Create(policy)
}

func (p *policyService) Delete(username, policyName string) error {
	return p.repo.PolicyRepo().Delete(username, policyName)
}

func (p *policyService) Update(policy *model.Policy) error {
	return p.repo.PolicyRepo().Update(policy)
}

func (p *policyService) Get(username, policyName string) (*model.Policy, error) {
	return p.repo.PolicyRepo().Get(username, policyName)
}

func (p *policyService) List(username string) (*model.PolicyList, error) {
	return p.repo.PolicyRepo().List(username)
}
