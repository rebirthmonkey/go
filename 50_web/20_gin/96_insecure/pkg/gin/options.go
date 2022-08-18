package gin

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
)

type Options struct {
	ServerOptions   *ServerOptions   `json:"server"   mapstructure:"server"`
	InsecureOptions *InsecureOptions `json:"insecure" mapstructure:"insecure"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	opt := Options{
		ServerOptions:   NewServerOptions(),
		InsecureOptions: NewInsecureOptions(),
	}

	return &opt
}

// Complete set default Options.
func (o *Options) Complete() error {
	return nil
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.ServerOptions.Validate()...)
	errs = append(errs, o.InsecureOptions.Validate()...)

	return errs
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	o.ServerOptions.ApplyTo(c)
	o.InsecureOptions.ApplyTo(c)
	
	return nil
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.ServerOptions.AddFlags(fss.FlagSet("gin server"))
	o.InsecureOptions.AddFlags(fss.FlagSet("insecure server"))

	return fss
}
