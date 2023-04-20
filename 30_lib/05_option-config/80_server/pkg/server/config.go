package server

// Config is a structure used to configure a Server.
// Its members are sorted roughly in order of importance for composers.
type Config struct {
	Healthz bool
	Mode    string
}

// CompletedConfig is the completed configuration for Server.
type CompletedConfig struct {
	*Config
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		Healthz: true,
		Mode:    "",
	}
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{c}
}

// New returns a new instance of Server from the given config.
func (c CompletedConfig) New() (*Server, error) {
	s := &Server{
		healthz: c.Healthz,
		mode:    c.Mode,
	}

	s.init()

	return s, nil
}
