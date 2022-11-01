// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo

// Repo defines the storage interface set to combine multiple resource repos.
type Repo interface {
	SecretRepo() SecretRepo
	Close() error
}

var client Repo

// Client return the store client instance.
func Client() Repo {
	return client
}

// SetClient set the store client.
func SetClient(c Repo) {
	client = c
}
