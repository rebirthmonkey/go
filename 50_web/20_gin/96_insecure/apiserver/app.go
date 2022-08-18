package apiserver

import (
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/apiserver/server"
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/app"
)

func NewApp(basename string) *app.App {
	opts := server.NewOptions()
	application := app.NewApp("demo app",
		basename,
		app.WithOptions(opts),
		app.WithDescription("demo app description"),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *server.Options) app.RunFunc {
	return func(basename string) error {
		cfg, err := server.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		server, err := server.CreateServer(cfg)
		if err != nil {
			return err
		}

		return server.PrepareRun().Run()
	}
}
