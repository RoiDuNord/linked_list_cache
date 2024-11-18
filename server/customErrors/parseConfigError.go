package customerrors

import "fmt"

type ConfigError struct {
	ErrorText string `json:"error"`
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf(e.ErrorText)
}

func ParseConfigError(err string) *ConfigError {
	return &ConfigError{ErrorText: err}
}
