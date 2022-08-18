package config

import (
	"github.com/rebirthmonkey/go/50_web/.bak/95_insecure/apiserver/options"
)

// Config is the running configuration structure of the app.
type Config struct {
	*options.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line or configuration file option.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
