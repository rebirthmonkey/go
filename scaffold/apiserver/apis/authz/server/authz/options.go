package authz

import (
	"github.com/spf13/pflag"
)

type Options struct {
	PolicyServer string `json:"policy-server" mapstructure:"policy-server"`
}

func NewOptions() *Options {
	return &Options{
		//PolicyServer: "127.0.0.1:30080",
	}
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *Options) Validate() []error {
	var errors []error
	return errors
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *Config) error {
	c.PolicyServer = o.PolicyServer
	return nil
}

// AddFlags adds flags related to features for a specific api server to the
// specified FlagSet.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
}
