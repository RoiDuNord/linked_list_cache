package parametervaluetointslice

import (
	"sort"
	valstostrs "server/server/convertingResponseToIntSlice/convertingFuncs/parameterValueToIntSlice/parameterValueToStringSlice"
	strstoints "server/server/convertingResponseToIntSlice/convertingFuncs/parameterValueToIntSlice/stringSliceToIntSlice"
)

func ParameterValueToSortedIntSlice(parameterValue string) ([]int, error) {

	stringSlice, err := valstostrs.ParameterValueToStringSlice(parameterValue)
	if err != nil {
		return nil, err
	}

	intSlice, err := strstoints.StringSliceToIntSlice(parameterValue, stringSlice)
	if err != nil {
		return nil, err
	}

	sort.Ints(intSlice)
	return intSlice, nil
}
