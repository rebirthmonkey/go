package apiserver

import (
	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/apiserver/config"
	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/apiserver/options"
	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/pkg/app"
)

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("demo APIServer",
		basename,
		app.WithOptions(opts),
		app.WithDescription("demo APIServer description"),
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
