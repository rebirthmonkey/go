package server

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_server/pkg/server"
)

type Options struct {
	ServerOptions *server.Options `json:"server"   mapstructure:"server"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		ServerOptions: server.NewOptions(),
	}
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.ServerOptions.Validate()...)

	return errs
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	o.ServerOptions.ApplyTo(c.GenericConfig)

	return nil
}

// Flags returns flags for a specific Server by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.ServerOptions.AddFlags(fss.FlagSet("server"))
	return fss
}
