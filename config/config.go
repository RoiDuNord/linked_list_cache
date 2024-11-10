package config

type Config struct {
	Port     int
	Factor   int
	LogLevel string
}

func NewConfig() *Config {
	return &Config{}
}
