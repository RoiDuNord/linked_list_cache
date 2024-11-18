package config

import (
	"fmt"
)

func (c *Config) validate() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port number (%d)", c.Port)
	}
	if c.Factor < 0 {
		return fmt.Errorf("factor must be greater than or equal to zero (%d)", c.Factor)
	}
	return nil
}
