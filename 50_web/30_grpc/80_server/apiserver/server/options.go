package server

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/gin"
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/grpc"
)

type Options struct {
	GinOptions  *gin.Options  `json:"gin"   mapstructure:"gin"`
	GrpcOptions *grpc.Options `json:"grpc"   mapstructure:"grpc"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		GinOptions:  gin.NewOptions(),
		GrpcOptions: grpc.NewOptions(),
	}
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GinOptions.Validate()...)
	errs = append(errs, o.GrpcOptions.Validate()...)

	return errs
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	err := o.GinOptions.ApplyTo(c.GinConfig)
	err = o.GrpcOptions.ApplyTo(c.GrpcConfig)

	return err
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GinOptions.AddFlags()
	o.GrpcOptions.AddFlags(fss.FlagSet("grpc"))

	return fss
}
