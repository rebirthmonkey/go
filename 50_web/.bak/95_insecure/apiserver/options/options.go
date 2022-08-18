package options

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/rebirthmonkey/go/50_web/.bak/95_insecure/pkg/options"
	"github.com/rebirthmonkey/go/50_web/.bak/95_insecure/pkg/server"
)

type Options struct {
	GenericServerRunOptions *options.ServerRunOptions       `json:"server"   mapstructure:"server"`
	InsecureServing         *options.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	opt := Options{
		GenericServerRunOptions: options.NewServerRunOptions(),
		InsecureServing:         options.NewInsecureServingOptions(),
	}

	return &opt
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))

	return fss
}

// Complete set default Options.
func (o *Options) Complete() error {
	return nil
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GenericServerRunOptions.Validate()...)

	return errs
}
