package mysql

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

	s := &Server{
		Host:     c.Host,
		Username: c.Username,
		Password: c.Password,
		Database: c.Database,
		//db:       db,
	}

	s.init()

	return s, nil
}
