package server

import (
	"github.com/rebirthmonkey/go/50_web/20_gin/97_secure/pkg/gin"
)

type Config struct {
	GinConfig *gin.Config
}

type CompletedConfig struct {
	CompletedGinConfig *gin.CompletedConfig
}

func NewConfig() *Config {
	return &Config{
		GinConfig: gin.NewConfig(),
	}
}

func (c *Config) Complete() *CompletedConfig {
	completedGinConfig := c.GinConfig.Complete()

	return &CompletedConfig{completedGinConfig}
}

func (c *CompletedConfig) New() (*Server, error) {
	ginServer, err := c.CompletedGinConfig.New()
	if err != nil {
		return nil, err
	}

	server := &Server{
		ginServer: ginServer,
	}

	return server, nil
}
