package apiserver

import (
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/server"
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/app"
)

func NewApp(basename string) *app.App {
	opts := server.NewOptions()
	application := app.NewApp("apiserver",
		basename,
		app.WithOptions(opts),
		app.WithDescription("apiserver description"),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *server.Options) app.RunFunc {
	return func(basename string) error {
		server, err := server.NewServer(opts)
		if err != nil {
			return err
		}

		return server.PrepareRun().Run()
	}
}
