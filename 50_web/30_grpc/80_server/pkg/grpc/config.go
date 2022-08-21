package grpc

import (
	"fmt"
	"google.golang.org/grpc"
)

type Config struct {
	Address string
}

type CompletedConfig struct {
	*Config
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Complete() *CompletedConfig {
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxx grpc config Address is:", c.Address)

	return &CompletedConfig{c}
}

func (c CompletedConfig) New() (*Server, error) {

	s := &Server{
		Address: c.Address,
		Server:  grpc.NewServer(),
	}

	s.init()

	return s, nil
}
