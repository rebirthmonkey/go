package options

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	genericoptions "github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/pkg/options"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions `json:"server"   mapstructure:"server"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	opt := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
	}

	return &opt
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))

	return fss
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GenericServerRunOptions.Validate()...)

	return errs
}
