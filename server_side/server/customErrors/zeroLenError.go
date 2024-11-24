package customerrors

import (
	"fmt"
)

type ZeroLen struct {
	ParameterValue string
}

func (e *ZeroLen) Error() string {
	return fmt.Sprintf("zero length error for parameter value: %s", e.ParameterValue)
}

func ZeroLenError(parameterValue string) *ZeroLen {
	return &ZeroLen{ParameterValue: parameterValue}
}
