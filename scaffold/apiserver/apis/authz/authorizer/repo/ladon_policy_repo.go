// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo

import (
	"github.com/ory/ladon"
)

// LadonPolicyRepo defines the ladon policy interface.
type LadonPolicyRepo interface {
	List() ([]*ladon.DefaultPolicy, error)
	//ListByUsername(string) ([]*ladon.DefaultPolicy, error)
	//SetList(string, []*ladon.DefaultPolicy)
	//Clear()
}
