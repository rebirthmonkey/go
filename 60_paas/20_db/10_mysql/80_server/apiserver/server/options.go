package server

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/grpc"
	"github.com/rebirthmonkey/go/pkg/mysql"
	"sync"
)

type Options struct {
	MysqlOptions *mysql.Options `json:"mysql"   mapstructure:"mysql"`
	GinOptions   *gin.Options   `json:"gin"   mapstructure:"gin"`
	GrpcOptions  *grpc.Options  `json:"grpc"   mapstructure:"grpc"`
}

var (
	opt  Options
	once sync.Once
)

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	once.Do(func() {
		opt = Options{
			MysqlOptions: mysql.NewOptions(),
			GinOptions:   gin.NewOptions(),
			GrpcOptions:  grpc.NewOptions(),
		}
	})

	return &opt
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.MysqlOptions.Validate()...)
	errs = append(errs, o.GinOptions.Validate()...)
	errs = append(errs, o.GrpcOptions.Validate()...)

	return errs
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	err := o.MysqlOptions.ApplyTo(c.MysqlConfig)
	err = o.GinOptions.ApplyTo(c.GinConfig)
	err = o.GrpcOptions.ApplyTo(c.GrpcConfig)

	return err
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GinOptions.AddFlags()
	o.GrpcOptions.AddFlags(fss.FlagSet("grpc"))
	o.MysqlOptions.AddFlags(fss.FlagSet("msyql"))

	return fss
}
