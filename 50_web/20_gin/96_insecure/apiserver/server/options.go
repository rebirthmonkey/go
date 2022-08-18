package server

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/gin"
)

type Options struct {
	GinOptions *gin.Options
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	opt := Options{
		GinOptions: gin.NewOptions(),
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

	errs = append(errs, o.GinOptions.Validate()...)

	return errs
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	o.GinOptions.ApplyTo(c.GinConfig)

	return nil
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GinOptions.Flags()

	return fss
}
