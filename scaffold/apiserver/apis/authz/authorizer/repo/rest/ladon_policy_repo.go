// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/ory/ladon"
	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/log"

	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/model/v1"
	authzRepo "github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/authorizer/repo"
)

type ladonPolicyRepo struct {
	policyRestURL string
}

var (
	landonPolicyRepoInstance *ladonPolicyRepo
	onceCache                sync.Once
)

var _ authzRepo.LadonPolicyRepo = (*ladonPolicyRepo)(nil)

func newLadonPolicyRepo(cfg *gin.CompletedConfig) authzRepo.LadonPolicyRepo {
	onceCache.Do(func() {
		landonPolicyRepoInstance = &ladonPolicyRepo{
			policyRestURL: "http://" + cfg.Insecure.Address + "/v1/policies",
		}
	})

	return landonPolicyRepoInstance
}

func (p *ladonPolicyRepo) List() ([]*ladon.DefaultPolicy, error) {
	log.Info("[Authorizer] PolicyRepo: list")

	response, err := http.Get(p.policyRestURL)
	if err != nil {
		log.Error("[AuthzPolicy/repo/rest] List: HTTP Response error")
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("[AuthzPolicy/repo/rest] List: HTTP Body error")
		return nil, err
	}

	var policyList model.PolicyList

	err = json.Unmarshal(body, &policyList)
	if err != nil {
		log.Error("[AuthzPolicy/repo/rest] List: HTTP JSON error")
		return nil, err
	}

	var ladonPolicyList []*ladon.DefaultPolicy

	for _, v := range policyList.Items {
		ladonPolicyList = append(ladonPolicyList, &v.AuthzPolicy.DefaultPolicy)
	}

	return ladonPolicyList, nil
}
