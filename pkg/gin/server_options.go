package gin

import (
	"github.com/spf13/pflag"
)

// ServerOptions contains the options while running a generic api server.
type ServerOptions struct {
	Mode        string   `json:"mode"        mapstructure:"mode"`
	Healthz     bool     `json:"healthz"     mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

// NewServerOptions creates a new ServerOptions object with default parameters.
func NewServerOptions() *ServerOptions {
	defaults := NewConfig()

	return &ServerOptions{
		Mode:        defaults.Mode,
		Healthz:     defaults.Healthz,
		Middlewares: defaults.Middlewares,
	}
}

// Validate checks validation of ServerRunOptions.
func (o *ServerOptions) Validate() []error {
	errors := []error{}

	return errors
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *ServerOptions) ApplyTo(c *Config) error {
	c.Mode = o.Mode
	c.Middlewares = o.Middlewares
	c.Healthz = o.Healthz

	return nil
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet.
func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs.StringVar(&o.Mode, "server.mode", o.Mode, ""+
		"Start the server in a specified server mode. Supported server mode: debug, test, release.")

	fs.StringSliceVar(&o.Middlewares, "server.middlewares", o.Middlewares, ""+
		"List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.")

	fs.BoolVar(&o.Healthz, "server.healthz", o.Healthz, ""+
		"Add self readiness check and install /healthz router.")
}
