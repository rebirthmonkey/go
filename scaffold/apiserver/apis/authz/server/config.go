// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"sync"

	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/grpc"
	"github.com/rebirthmonkey/go/pkg/log"
	"github.com/rebirthmonkey/go/pkg/mysql"
)

// Config is the running configuration structure of the server.
type Config struct {
	LogConfig       *log.Config
	MysqlConfig     *mysql.Config
	GinConfig       *gin.Config
	GrpcConfig      *grpc.Config
	ApiserverConfig *gin.Config
}

// CompletedConfig is the complete configuration structure of the server.
type CompletedConfig struct {
	CompletedLogConfig       *log.CompletedConfig
	CompletedMysqlConfig     *mysql.CompletedConfig
	CompletedGinConfig       *gin.CompletedConfig
	CompletedGrpcConfig      *grpc.CompletedConfig
	CompletedApiserverConfig *gin.CompletedConfig
}

var (
	config     CompletedConfig
	onceConfig sync.Once
)

// NewConfig creates a running configuration instance based
// on a given command line or configuration file option.
func NewConfig() *Config {
	return &Config{
		LogConfig:       log.NewConfig(),
		MysqlConfig:     mysql.NewConfig(),
		GinConfig:       gin.NewConfig(),
		GrpcConfig:      grpc.NewConfig(),
		ApiserverConfig: gin.NewConfig(),
	}
}

// Complete set default Configs.
func (c *Config) Complete() *CompletedConfig {

	onceConfig.Do(func() {
		config = CompletedConfig{
			CompletedLogConfig:       c.LogConfig.Complete(),
			CompletedMysqlConfig:     c.MysqlConfig.Complete(),
			CompletedGinConfig:       c.GinConfig.Complete(),
			CompletedGrpcConfig:      c.GrpcConfig.Complete(),
			CompletedApiserverConfig: c.ApiserverConfig.Complete(),
		}
	})

	return &config
}

// New creates a new server based on the configuration
func (c *CompletedConfig) New() (*Server, error) {
	err := c.CompletedLogConfig.New()
	if err != nil {
		log.Fatalf("Failed to launch Log: %s", err.Error())
		return nil, err
	}

	ginServer, err := c.CompletedGinConfig.New()
	if err != nil {
		log.Fatalf("Failed to launch Gin server: %s", err.Error())
		return nil, err
	}

	grpcServer, err := c.CompletedGrpcConfig.New()
	if err != nil {
		log.Fatalf("Failed to launch Grpc Server: %s", err.Error())
		return nil, err
	}

	server := &Server{
		ginServer:  ginServer,
		grpcServer: grpcServer,
	}

	return server, nil
}
