package grpc

import (
	"fmt"
	"github.com/spf13/pflag"
	"net"
	"strconv"
)

type Options struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

func NewOptions() *Options {
	return &Options{}
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errors []error

	if o.BindPort < 0 || o.BindPort > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--grpc.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off port",
				o.BindPort,
			),
		)
	}

	return errors
}

func (o *Options) ApplyTo(c *Config) error {
	c.Address = net.JoinHostPort(o.BindAddress, strconv.Itoa(o.BindPort))

	return nil
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.BindAddress, "grpc.bind-address", o.BindAddress, ""+
		"The IP address on which to serve the --grpc.bind-port "+
		"(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")

	fs.IntVar(&o.BindPort, "grpc.bind-port", o.BindPort, ""+
		"The port on which to serve unsecured, unauthenticated access. It is assumed "+
		"that firewall rules are set up such that this port is not reachable from outside of "+
		"the deployed machine and that port 8081 on the iam public address is proxied to this "+
		"port. This is performed by nginx in the default setup. Set to zero to disable.")
}
