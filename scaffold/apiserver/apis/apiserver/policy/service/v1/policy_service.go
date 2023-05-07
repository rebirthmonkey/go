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
	Delete(policyName string) error
	Update(policy *model.Policy) error
	Get(policyName string) (*model.Policy, error)
	List() (*model.PolicyList, error)
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

func (p *policyService) Delete(policyName string) error {
	return p.repo.PolicyRepo().Delete(policyName)
}

func (p *policyService) Update(policy *model.Policy) error {
	return p.repo.PolicyRepo().Update(policy)
}

func (p *policyService) Get(policyName string) (*model.Policy, error) {
	return p.repo.PolicyRepo().Get(policyName)
}

func (p *policyService) List() (*model.PolicyList, error) {
	return p.repo.PolicyRepo().List()
}
