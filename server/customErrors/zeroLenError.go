package customerrors

import (
	"fmt"
)

type ZeroLenError struct {
	parameterValue string
}

func (e *ZeroLenError) Error() string {
	return fmt.Sprintf("zero length error for parameter value: %s", e.parameterValue)
}

func NewZeroLenError(parameterValue string) *ZeroLenError {
	return &ZeroLenError{parameterValue: parameterValue}
}
