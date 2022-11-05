// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/rebirthmonkey/go/pkg/app"
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/server"
)

// NewApp creates an App object with default parameters.
func NewApp(basename string) *app.App {
	opts := server.NewOptions()
	application := app.NewApp("authz",
		basename,
		app.WithOptions(opts),
		app.WithDescription("authz description"),
		app.WithRunFunc(run(opts)),
	)

	return application
}

// run launches the App object.
func run(opts *server.Options) app.RunFunc {
	return func(basename string) error {
		server, err := server.NewServer(opts)
		if err != nil {
			return err
		}

		return server.PrepareRun().Run()
	}
}
