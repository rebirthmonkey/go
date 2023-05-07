// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	model "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/secret/model/v1"
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/secret/repo"
)

type SecretService interface {
	Create(secret *model.Secret) error
	Update(secret *model.Secret) error
	Delete(secretID string) error
	Get(secretID string) (*model.Secret, error)
	List() (*model.SecretList, error)
}

type secretService struct {
	repo repo.Repo
}

var _ SecretService = (*secretService)(nil)

// newSecretService creates and returns the secret service instance.
func newSecretService(repo repo.Repo) SecretService {
	return &secretService{repo: repo}
}

func (s *secretService) Create(secret *model.Secret) error {
	return s.repo.SecretRepo().Create(secret)
}

func (s *secretService) Delete(secretName string) error {
	return s.repo.SecretRepo().Delete(secretName)
}

func (s *secretService) Update(secret *model.Secret) error {
	return s.repo.SecretRepo().Update(secret)
}

func (s *secretService) Get(secretName string) (*model.Secret, error) {
	return s.repo.SecretRepo().Get(secretName)
}

func (s *secretService) List() (*model.SecretList, error) {
	return s.repo.SecretRepo().List()
}
