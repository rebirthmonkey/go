// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo

import (
	model "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/secret/model/v1"
)

// SecretRepo defines the secret resources.
type SecretRepo interface {
	Create(secret *model.Secret) error
	Delete(username, secretName string) error
	DeleteByUser(username string) error
	Update(secret *model.Secret) error
	Get(username, secretName string) (*model.Secret, error)
	List(username string) (*model.SecretList, error)
}
