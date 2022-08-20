// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo

type Repo interface {
	GetUserRepo() UserRepo
	Close() error
}

var repo Repo

func GetRepo() Repo {
	return repo
}

func SetRepo(r Repo) {
	repo = r
}
