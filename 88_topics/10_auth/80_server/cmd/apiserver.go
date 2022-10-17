// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver"
)

func main() {
	apiserver.NewApp("apiserver").Run()
}
