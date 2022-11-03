// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo

import (
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/model/v1"
)

// PolicyRepo defines the secret resources.
type PolicyRepo interface {
	Create(policy *model.Policy) error
	Delete(username string, policyName string) error
	DeleteByUser(username string) error
	Update(policy *model.Policy) error
	Get(username string, policyName string) (*model.Policy, error)
	List(username string) (*model.PolicyList, error)
}
