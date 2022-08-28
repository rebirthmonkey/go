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

type Config struct {
	LogConfig   *log.Config
	MysqlConfig *mysql.Config
	GinConfig   *gin.Config
	GrpcConfig  *grpc.Config
}

type CompletedConfig struct {
	CompletedLogConfig   *log.CompletedConfig
	CompletedMysqlConfig *mysql.CompletedConfig
	CompletedGinConfig   *gin.CompletedConfig
	CompletedGrpcConfig  *grpc.CompletedConfig
}

var (
	config     CompletedConfig
	onceConfig sync.Once
)

func NewConfig() *Config {
	return &Config{
		LogConfig:   log.NewConfig(),
		MysqlConfig: mysql.NewConfig(),
		GinConfig:   gin.NewConfig(),
		GrpcConfig:  grpc.NewConfig(),
	}
}

func (c *Config) Complete() *CompletedConfig {

	onceConfig.Do(func() {
		config = CompletedConfig{
			CompletedLogConfig:   c.LogConfig.Complete(),
			CompletedMysqlConfig: c.MysqlConfig.Complete(),
			CompletedGinConfig:   c.GinConfig.Complete(),
			CompletedGrpcConfig:  c.GrpcConfig.Complete(),
		}
	})

	return &config
}

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
