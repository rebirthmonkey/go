package server

import (
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/gin"
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/grpc"
)

type Config struct {
	GinConfig  *gin.Config
	GrpcConfig *grpc.Config
}

type CompletedConfig struct {
	CompletedGinConfig  *gin.CompletedConfig
	CompletedGrpcConfig *grpc.CompletedConfig
}

func NewConfig() *Config {
	return &Config{
		GinConfig:  gin.NewConfig(),
		GrpcConfig: grpc.NewConfig(),
	}
}

func (c *Config) Complete() *CompletedConfig {
	completedGinConfig := c.GinConfig.Complete()
	completedGrpcConfig := c.GrpcConfig.Complete()

	return &CompletedConfig{
		CompletedGinConfig:  completedGinConfig,
		CompletedGrpcConfig: completedGrpcConfig,
	}
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
