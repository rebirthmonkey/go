package gin

import (
	"github.com/gin-gonic/gin"
)

// InsecureConfig holds configuration of the insecure http server.
type InsecureConfig struct {
	Address string
}

// SecureConfig holds configuration of the secure https server.
type SecureConfig struct {
	Address  string
	CertFile string
	KeyFile  string
}

// Config is a structure used to configure a GinServer.
// Its members are sorted roughly in order of importance for composers.
type Config struct {
	Mode        string
	Middlewares []string
	Healthz     bool

	Insecure *InsecureConfig
	Secure   *SecureConfig
}

// CompletedConfig is the completed configuration for GenericAPIServer.
type CompletedConfig struct {
	*Config
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		Healthz:     true,
		Mode:        gin.ReleaseMode,
		Middlewares: []string{},
	}
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{c}
}

// New returns a new instance of GenericAPIServer from the given config.
func (c CompletedConfig) New() (*Server, error) {

	gin.SetMode(c.Mode)

	s := &Server{
		Healthz:     c.Healthz,
		Middlewares: c.Middlewares,

		Engine: gin.New(),

		Insecure: c.Insecure,
		Secure:   c.Secure,
	}

	s.init()

	return s, nil
}
