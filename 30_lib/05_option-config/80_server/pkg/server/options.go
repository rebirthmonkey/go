package server

import (
	"github.com/spf13/pflag"
)

type Options struct {
	Healthz bool   `json:"healthz"     mapstructure:"healthz"`
	Mode    string `json:"mode"     mapstructure:"mode"`
}

func NewOptions() *Options {
	return &Options{
		Healthz: true,
		Mode:    "",
	}
}

func (o *Options) Validate() []error {
	errors := []error{}

	return errors
}

func (o *Options) ApplyTo(c *Config) error {
	c.Healthz = o.Healthz
	c.Mode = o.Mode
	return nil
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&o.Healthz, "server.healthz", o.Healthz, "healthz.")
	fs.StringVar(&o.Mode, "server.mode", o.Mode, "Mode.")
}
