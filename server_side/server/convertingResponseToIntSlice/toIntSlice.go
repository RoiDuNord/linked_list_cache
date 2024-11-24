package converting

import (
	"fmt"
	valstoints "server/server/convertingResponseToIntSlice/convertingFuncs/parameterValueToIntSlice"
)

type GetInfoRequestData string

func (parameterValue GetInfoRequestData) ToIntSlice(parameter string) ([]int, error) {
	parameterValueString := string(parameterValue)
	if err := isParameterNumbers(parameter); err != nil {
		return nil, err
	}
	slice, err := valstoints.ParameterValueToSortedIntSlice(parameterValueString)
	return slice, err
}

func isParameterNumbers(parameter string) error {
	if parameter != "numbers" {
		return fmt.Errorf("server doesn't contain information about requested parameter '%s'", parameter)
	}
	return nil
}
