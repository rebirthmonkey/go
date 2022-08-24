package server

import (
	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/grpc"
	"github.com/rebirthmonkey/go/pkg/mysql"
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

func NewConfig() *Config {
	return &Config{
		MysqlConfig: mysql.NewConfig(),
		GinConfig:   gin.NewConfig(),
		GrpcConfig:  grpc.NewConfig(),
	}
}

func (c *Config) Complete() *CompletedConfig {
	completedMysqlConfig := c.MysqlConfig.Complete()
	completedGinConfig := c.GinConfig.Complete()
	completedGrpcConfig := c.GrpcConfig.Complete()

	return &CompletedConfig{
		CompletedMysqlConfig: completedMysqlConfig,
		CompletedGinConfig:   completedGinConfig,
		CompletedGrpcConfig:  completedGrpcConfig,
	}
}

func (c *CompletedConfig) New() (*Server, error) {
	mysqlServer, err := c.CompletedMysqlConfig.New()
	if err != nil {
		return nil, err
	}

	ginServer, err := c.CompletedGinConfig.New()
	if err != nil {
		return nil, err
	}

	grpcServer, err := c.CompletedGrpcConfig.New()
	if err != nil {
		return nil, err
	}

	server := &Server{
		mysqlServer: mysqlServer,
		ginServer:   ginServer,
		grpcServer:  grpcServer,
	}

	return server, nil
}
