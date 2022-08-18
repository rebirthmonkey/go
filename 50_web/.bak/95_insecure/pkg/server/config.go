package server

import "github.com/gin-gonic/gin"

// Config is a structure used to configure a GenericAPIServer.
// Its members are sorted roughly in order of importance for composers.
type Config struct {
	InsecureServing *InsecureServingInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
}

// InsecureServingInfo holds configuration of the insecure http server.
type InsecureServingInfo struct {
	Address string
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		Healthz:     true,
		Mode:        gin.ReleaseMode,
		Middlewares: []string{},
	}
}

// CompletedConfig is the completed configuration for GenericAPIServer.
type CompletedConfig struct {
	*Config
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

// New returns a new instance of GenericAPIServer from the given config.
func (c CompletedConfig) New() (*GenericAPIServer, error) {

	gin.SetMode(c.Mode)

	s := &GenericAPIServer{
		InsecureServingInfo: c.InsecureServing,
		healthz:             c.Healthz,
		middlewares:         c.Middlewares,
		Engine:              gin.New(),
	}

	initGenericAPIServer(s)

	return s, nil
}
