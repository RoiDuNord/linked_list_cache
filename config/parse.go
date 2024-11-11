package config

import (
	"fmt"
)

func ParseConfig() (*Config, error) {
	file := "/Volumes/Data/workWithCache/config/config.yaml"
	cfg, err := loadConfig(file)
	if err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	fmt.Printf("Server will start on port: %d\n", cfg.Port)
	fmt.Printf("Factor: %d\n", cfg.Factor)
	fmt.Printf("Log Level: %s\n", cfg.LogLevel)

	fmt.Print(cfg)

	return cfg, nil
}
