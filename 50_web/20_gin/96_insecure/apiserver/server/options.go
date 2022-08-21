package server

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/rebirthmonkey/go/50_web/20_gin/96_insecure/pkg/gin"
)

type Options struct {
	GinOptions *gin.Options `json:"gin"   mapstructure:"gin"`
}

func NewOptions() *Options {
	return &Options{
		GinOptions: gin.NewOptions(),
	}
}

func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GinOptions.Validate()...)

	return errs
}

func (o *Options) ApplyTo(c *Config) error {
	err := o.GinOptions.ApplyTo(c.GinConfig)

	return err
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GinOptions.AddFlags()

	return fss
}
