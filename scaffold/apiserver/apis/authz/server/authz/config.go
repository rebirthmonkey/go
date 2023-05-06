package authz

// Config holds configuration of the insecure http server.
type Config struct {
	PolicyServer string
}

// CompletedConfig is the completed configuration for GenericAPIServer.
type CompletedConfig struct {
	*Config
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		PolicyServer: "",
	}
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{c}
}

func (c CompletedConfig) New() error {

	return nil
}
