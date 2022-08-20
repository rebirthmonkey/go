package server

import (
	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/pkg/server"
)

type Config struct {
	GenericConfig *server.Config
}

type CompletedConfig struct {
	CompletedGenericConfig *server.CompletedConfig
}

func NewConfig() *Config {
	return &Config{
		GenericConfig: server.NewConfig(),
	}
}

func (c *Config) Complete() *CompletedConfig {
	completedGenericConfig := c.GenericConfig.Complete()

	return &CompletedConfig{completedGenericConfig}
}

func (c *CompletedConfig) New() (*Server, error) {
	genericServer, err := c.CompletedGenericConfig.New()
	if err != nil {
		return nil, err
	}

	server := &Server{
		genericServer: genericServer,
	}

	return server, nil
}
