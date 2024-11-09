package converting

import (
	"fmt"
	convfuncs "workWithCache/server/elements/converting/convertingFuncs/parameterValueToIntSlice"
)

type GetInfoRequestData string

func (parameterValue GetInfoRequestData) ToIntSlice(parameter string) ([]int, error) {
	parameterValueString := string(parameterValue)
	if err := isParameterNumbers(parameter); err != nil {
		return nil, err
	}
	slice, err := convfuncs.ParameterValueToSortedIntSlice(parameterValueString)
	return slice, err
}

func isParameterNumbers(parameter string) error {
	if parameter != "numbers" {
		return fmt.Errorf("server doesn't contain information for requested parameter '%s'", parameter)
	}
	return nil
}
