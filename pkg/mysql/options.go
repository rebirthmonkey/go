package mysql

import (
	"fmt"

	"github.com/spf13/pflag"
)

type Options struct {
	Host     string `json:"host" mapstructure:"host"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

func NewOptions() *Options {
	return &Options{
		Host:     "",
		Username: "",
		Password: "",
		Database: "",
	}
}

func (o *Options) Validate() []error {
	var errors []error

	if o.Host == "" || o.Username == "" || o.Database == "" {
		errors = append(errors, fmt.Errorf("invalide config"))
	}

	return errors
}

func (o *Options) ApplyTo(c *Config) error {
	c.Host = o.Host
	c.Username = o.Username
	c.Password = o.Password
	c.Database = o.Database

	return nil
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql.host", o.Host, "MySQL Host.")
}
