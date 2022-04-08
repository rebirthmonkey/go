// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/rebirthmonkey/pkg/log"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

type Options struct{
	Server struct{
		Mode string `mapstructure:"mode"`
		Healthz string `mapstructure:"healthz"`
	}
	Insecure struct {
		Address string	`mapstructure:"bind-address"`
		Port string		`mapstructure:"bind-port"`
	}
	Secure struct {
		Address string	`mapstructure:"bind-address"`
		Port string		`mapstructure:"bind-port"`
		TLS struct {
			CertFile string `mapstructure:"cert-file"`
			PrivateKeyFile string `mapstructure:"private-key-file"`
		}
	}
	Log struct {
		Level string `mapstructure:"level"`
		Format string `mapstructure:"format"`
		EnableColor string `mapstructure:"enable-color"`
		DisableCaller string `mapstructure:"disable-caller"`
		OutputPaths string `mapstructure:"output-paths"`
		ErrorOutputPaths string `mapstructure:"error-output-paths"`
	}
	GRPC struct {
		Address string	`mapstructure:"bind-address"`
		Port string		`mapstructure:"bind-port"`
	}
}

var options Options

func (o *Options) load(cfgFileName string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(cfgFileName)
	}
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("failed to read configuration file(%s): %v\n", cfgFile, err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&options); err != nil {
		log.Panicf("unable to decode into struct: %s \n", err)
	}
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}

