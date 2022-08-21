package gin

import (
	"fmt"
	"github.com/spf13/pflag"
	"net"
	"strconv"
)

type SecureOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
	TLS         struct {
		CertFile       string `json:"cert-file" mapstructure:"cert-file"`
		PrivateKeyFile string `json:"private-key-file" mapstructure:"private-key-file"`
	}
}

// NewSecureOptions is for creating an unauthenticated, unauthorized, insecure port.
// No one should be using these anymore.
func NewSecureOptions() *SecureOptions {
	return &SecureOptions{
		//BindAddress: "127.0.0.1",
		//BindPort:    8080,
	}
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *SecureOptions) Validate() []error {
	var errors []error

	if o.BindPort < 0 || o.BindPort > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--insecure.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port",
				o.BindPort,
			),
		)
	}

	return errors
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *SecureOptions) ApplyTo(c *Config) error {
	c.Secure = &SecureConfig{
		Address:  net.JoinHostPort(o.BindAddress, strconv.Itoa(o.BindPort)),
		CertFile: o.TLS.CertFile,
		KeyFile:  o.TLS.PrivateKeyFile,
	}

	return nil
}

// AddFlags adds flags related to features for a specific api server to the
// specified FlagSet.
func (o *SecureOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.BindAddress, "secure.bind-address", o.BindAddress, ""+
		"The IP address on which to serve the --secure.bind-port "+
		"(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")

	fs.IntVar(&o.BindPort, "secure.bind-port", o.BindPort, ""+
		"The port on which to serve unsecured, unauthenticated access. It is assumed "+
		"that firewall rules are set up such that this port is not reachable from outside of "+
		"the deployed machine and that port 443 on the iam public address is proxied to this "+
		"port. This is performed by nginx in the default setup. Set to zero to disable.")
}
