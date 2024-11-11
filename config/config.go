package config

type Config struct {
	Port     int    `yaml:"port"`
	Factor   int    `yaml:"factor"`
	LogLevel string `yaml:"loglevel"`
}
