package gin

import (
	"github.com/spf13/pflag"
)

type ServerOptions struct {
	Mode        string   `json:"mode"        mapstructure:"mode"`
	Healthz     bool     `json:"healthz"     mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerOptions() *ServerOptions {
	//defaults := NewConfig()

	//return &ServerOptions{
	//	Mode:        defaults.Mode,
	//	Healthz:     defaults.Healthz,
	//	Middlewares: defaults.Middlewares,
	//}

	return &ServerOptions{
		Mode:        "",
		Healthz:     true,
		Middlewares: []string{},
	}
}

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

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Mode, "gin.server.mode", o.Mode, "server mode")

	fs.StringSliceVar(&o.Middlewares, "gin.server.middlewares", o.Middlewares, ""+
		"List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.")

	fs.BoolVar(&o.Healthz, "gin.server.healthz", o.Healthz, "healthz")
}
