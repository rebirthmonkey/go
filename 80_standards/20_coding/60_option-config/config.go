// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

type Config struct {
	*Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line or configuration file option.
func CreateConfigFromOptions(opts *Options) (*Config, error) {
	return &Config{opts}, nil
}
