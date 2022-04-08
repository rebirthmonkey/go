// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/rebirthmonkey/pkg/log"
	"strconv"
)

type Config struct {
	Log					*log.Options
}


func CreateConfig() *Config {
	return &Config{
		Log:			 	&log.Options{},
	}
}

func (cfg *Config) BuildConfigFromOptions(opts *Options) error {
	cfg.Log.Level = opts.Log.Level
	cfg.Log.Format = opts.Log.Format
	cfg.Log.EnableColor, _ = strconv.ParseBool(opts.Log.EnableColor)
	cfg.Log.DisableCaller, _ = strconv.ParseBool(opts.Log.DisableCaller)
	cfg.Log.OutputPaths = append(cfg.Log.OutputPaths, opts.Log.OutputPaths)
	cfg.Log.ErrorOutputPaths = append(cfg.Log.ErrorOutputPaths, opts.Log.ErrorOutputPaths)
	return nil
}
