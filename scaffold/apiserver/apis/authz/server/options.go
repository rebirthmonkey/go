// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"sync"

	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/grpc"
	"github.com/rebirthmonkey/go/pkg/log"
	"github.com/rebirthmonkey/go/pkg/mysql"

	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/authz/server/authz"
)

// Options is the options of a server.
type Options struct {
	LogOptions   *log.Options   `json:"log"   mapstructure:"log"`
	MysqlOptions *mysql.Options `json:"mysql"   mapstructure:"mysql"`
	GinOptions   *gin.Options   `json:"gin"   mapstructure:"gin"`
	GrpcOptions  *grpc.Options  `json:"grpc"   mapstructure:"grpc"`
	AuthzOptions *authz.Options `json:"authz"   mapstructure:"authz"`
}

var (
	opt  Options
	once sync.Once
)

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	once.Do(func() {
		opt = Options{
			LogOptions:   log.NewOptions(),
			MysqlOptions: mysql.NewOptions(),
			GinOptions:   gin.NewOptions(),
			GrpcOptions:  grpc.NewOptions(),
			AuthzOptions: authz.NewOptions(),
		}
	})

	return &opt
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.LogOptions.Validate()...)
	errs = append(errs, o.MysqlOptions.Validate()...)
	errs = append(errs, o.GinOptions.Validate()...)
	errs = append(errs, o.GrpcOptions.Validate()...)
	errs = append(errs, o.AuthzOptions.Validate()...)

	return errs
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	if err := o.LogOptions.ApplyTo(c.LogConfig); err != nil {
		log.Panic(err.Error())
	}

	if err := o.MysqlOptions.ApplyTo(c.MysqlConfig); err != nil {
		log.Panic(err.Error())
	}

	if err := o.GinOptions.ApplyTo(c.GinConfig); err != nil {
		log.Panic(err.Error())
	}

	if err := o.GrpcOptions.ApplyTo(c.GrpcConfig); err != nil {
		log.Panic(err.Error())
	}

	if err := o.AuthzOptions.ApplyTo(c.AuthzConfig); err != nil {
		log.Panic(err.Error())
	}

	return nil
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GinOptions.AddFlags(fss.FlagSet("gin"))
	o.GrpcOptions.AddFlags(fss.FlagSet("grpc"))
	o.MysqlOptions.AddFlags(fss.FlagSet("mysql"))
	o.LogOptions.AddFlags(fss.FlagSet("log"))
	o.AuthzOptions.AddFlags(fss.FlagSet("authz"))

	return fss
}
