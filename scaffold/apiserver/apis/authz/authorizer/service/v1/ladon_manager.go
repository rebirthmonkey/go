// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/marmotedu/errors"
	"github.com/ory/ladon"

	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
)

type ladonPolicyManager struct {
	repo repo.Repo
}

// NewLadonManager initializes a new ladonPolicyManager for given repo.
func NewLadonManager(repo repo.Repo) ladon.Manager {
	return &ladonPolicyManager{
		repo: repo,
	}
}

// Create persists the policy.
func (*ladonPolicyManager) Create(policy ladon.Policy) error {
	return nil
}

// Update updates an existing policy.
func (*ladonPolicyManager) Update(policy ladon.Policy) error {
	return nil
}

// Get retrieves a policy.
func (*ladonPolicyManager) Get(id string) (ladon.Policy, error) {
	return nil, nil
}

// Delete removes a policy.
func (*ladonPolicyManager) Delete(id string) error {
	return nil
}

// GetAll retrieves all policies.
func (*ladonPolicyManager) GetAll(limit, offset int64) (ladon.Policies, error) {
	return nil, nil
}

// FindRequestCandidates returns candidates that could match the request object. It either returns
// a set that exactly matches the request, or a superset of it. If an error occurs, it returns nil and
// the error.
func (m *ladonPolicyManager) FindRequestCandidates(r *ladon.Request) (ladon.Policies, error) {
	//var username string
	//
	//if user, ok := r.Context["username"].(string); ok {
	//	username = user
	//}

	policies, err := m.repo.LadonPolicyRepo().List()
	if err != nil {
		return nil, errors.Wrap(err, "list policies failed")
	}

	ret := make([]ladon.Policy, 0, len(policies))
	for _, policy := range policies {
		ret = append(ret, policy)
	}

	return ret, nil
}

// FindPoliciesForSubject returns policies that could match the subject. It either returns
// a set of policies that applies to the subject, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *ladonPolicyManager) FindPoliciesForSubject(subject string) (ladon.Policies, error) {
	return nil, nil
}

// FindPoliciesForResource returns policies that could match the resource. It either returns
// a set of policies that apply to the resource, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *ladonPolicyManager) FindPoliciesForResource(resource string) (ladon.Policies, error) {
	return nil, nil
}
