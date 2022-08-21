package gin

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
)

type Options struct {
	ServerOptions   *ServerOptions   `json:"server"   mapstructure:"server"`
	InsecureOptions *InsecureOptions `json:"insecure" mapstructure:"insecure"`
	SecureOptions   *SecureOptions   `json:"secure" mapstructure:"secure"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		ServerOptions:   NewServerOptions(),
		InsecureOptions: NewInsecureOptions(),
		SecureOptions:   NewSecureOptions(),
	}
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.ServerOptions.Validate()...)
	errs = append(errs, o.InsecureOptions.Validate()...)
	errs = append(errs, o.SecureOptions.Validate()...)

	return errs
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	o.ServerOptions.ApplyTo(c)
	o.InsecureOptions.ApplyTo(c)
	o.SecureOptions.ApplyTo(c)

	return nil
}

// AddFlags returns flags for a specific APIServer by section name.
func (o *Options) AddFlags() (fss cliflag.NamedFlagSets) {
	o.ServerOptions.AddFlags(fss.FlagSet("server"))
	o.InsecureOptions.AddFlags(fss.FlagSet("insecure"))
	o.SecureOptions.AddFlags(fss.FlagSet("secure"))

	return fss
}
