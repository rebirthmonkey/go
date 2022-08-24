package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Username string
	Password string
	Database string
}

type CompletedConfig struct {
	*Config
}

func NewConfig() *Config {
	return &Config{
		Host:     "",
		Username: "",
		Password: "",
		Database: "",
	}
}

func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{c}
}

func (c CompletedConfig) New() (*Server, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		c.Username,
		c.Password,
		c.Host,
		c.Database,
		true,
		"Local")

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	s := &Server{
		Host:     c.Host,
		Username: c.Username,
		Password: c.Password,
		Database: c.Database,
		db:       db,
	}

	s.init()

	return s, nil
}
