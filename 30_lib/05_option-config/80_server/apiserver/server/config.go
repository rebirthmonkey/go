package server

import (
	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_server/pkg/server"
)

type Config struct {
	GenericConfig *server.Config
	//MysqlConfig   *mysql.Config
}

type CompletedConfig struct {
	CompletedGenericConfig *server.CompletedConfig
	//CompletedMysqlConfig   *mysql.CompletedConfig
}

func NewConfig() *Config {
	return &Config{
		GenericConfig: server.NewConfig(),
		//MysqlConfig:   mysql.NewConfig(),
	}
}

func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{
		CompletedGenericConfig: c.GenericConfig.Complete(),
		//CompletedMysqlConfig:   c.MysqlConfig.Complete(),
	}
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
