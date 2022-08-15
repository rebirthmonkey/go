package config

import "github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/apiserver/options"

// Config is the running configuration structure of the IAM pump service.
type Config struct {
	*options.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given IAM pump command line or configuration file option.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
