package options

import (
	"github.com/spf13/pflag"

	"github.com/rebirthmonkey/go/30_lib/05_option-config/80_option-config/pkg/server"
)

// ServerRunOptions contains the options while running a generic api server.
type ServerRunOptions struct {
	Healthz bool `json:"healthz"     mapstructure:"healthz"`
}

// NewServerRunOptions creates a new ServerRunOptions object with default parameters.
func NewServerRunOptions() *ServerRunOptions {
	defaults := server.NewConfig()

	return &ServerRunOptions{
		Healthz: defaults.Healthz,
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *ServerRunOptions) ApplyTo(c *server.Config) error {
	c.Healthz = s.Healthz

	return nil
}

// Validate checks validation of ServerRunOptions.
func (s *ServerRunOptions) Validate() []error {
	errors := []error{}

	return errors
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet.
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {

	fs.BoolVar(&s.Healthz, "healthz", true, "Add self readiness check /healthz.")

}
