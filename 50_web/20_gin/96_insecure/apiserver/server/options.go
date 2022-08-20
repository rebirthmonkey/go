package server

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/gin"
)

type Options struct {
	GinOptions *gin.Options `json:"gin"   mapstructure:"gin"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		GinOptions: gin.NewOptions(),
	}
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GinOptions.Validate()...)

	return errs
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	err := o.GinOptions.ApplyTo(c.GinConfig)

	return err
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GinOptions.AddFlags()

	return fss
}
