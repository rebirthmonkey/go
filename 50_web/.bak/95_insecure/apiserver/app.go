package apiserver

import (
	"github.com/rebirthmonkey/go/50_web/.bak/95_insecure/apiserver/config"
	"github.com/rebirthmonkey/go/50_web/.bak/95_insecure/apiserver/options"
	"github.com/rebirthmonkey/go/50_web/.bak/95_insecure/pkg/app"
)

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("demo app",
		basename,
		app.WithOptions(opts),
		app.WithDescription("demo app description"),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		server, err := createAPIServer(cfg)
		if err != nil {
			return err
		}

		return server.PrepareRun().Run()
	}
}
