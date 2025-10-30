package config

type Config struct {
	Network string
}

func NewConfig() *Config {
	return &Config{}
}
