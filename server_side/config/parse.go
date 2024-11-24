package config

func ParseConfig() (Config, error) {
	cfg, err := loadConfigFromFile("config.yaml")
	if err != nil {
		return Config{}, err
	}

	if err := cfg.validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
