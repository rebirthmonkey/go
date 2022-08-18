package server

import (
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/gin"
)

// Config is the running configuration structure of the app.
//type Config struct {
//	*Options
//}

type Config struct {
	GinConfig *gin.Config
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line or configuration file option.
func CreateConfigFromOptions(opts *Options) (*Config, error) {
	ginConfig := gin.NewConfig()
	opts.GinOptions.ApplyTo(ginConfig)
	return &Config{ginConfig}, nil
}
