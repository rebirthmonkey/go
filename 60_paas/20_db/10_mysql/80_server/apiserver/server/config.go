package server

import (
	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/grpc"
	"github.com/rebirthmonkey/go/pkg/mysql"
	"sync"
)

type Config struct {
	MysqlConfig *mysql.Config
	GinConfig   *gin.Config
	GrpcConfig  *grpc.Config
}

type CompletedConfig struct {
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
		MysqlConfig: mysql.NewConfig(),
		GinConfig:   gin.NewConfig(),
		GrpcConfig:  grpc.NewConfig(),
	}
}

func (c *Config) Complete() *CompletedConfig {

	onceConfig.Do(func() {
		config = CompletedConfig{
			CompletedMysqlConfig: c.MysqlConfig.Complete(),
			CompletedGinConfig:   c.GinConfig.Complete(),
			CompletedGrpcConfig:  c.GrpcConfig.Complete(),
		}
	})

	return &config
}

func (c *CompletedConfig) New() (*Server, error) {
	ginServer, err := c.CompletedGinConfig.New()
	if err != nil {
		return nil, err
	}

	grpcServer, err := c.CompletedGrpcConfig.New()
	if err != nil {
		return nil, err
	}

	server := &Server{
		ginServer:  ginServer,
		grpcServer: grpcServer,
	}

	return server, nil
}
