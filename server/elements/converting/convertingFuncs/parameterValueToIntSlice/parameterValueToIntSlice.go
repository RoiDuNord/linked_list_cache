package convfuncs

import (
	"sort"
	converting "workWithCache/server/elements/converting/convertingFuncs/parameterValueToIntSlice/parameterValueToStringSlice"
	conv "workWithCache/server/elements/converting/convertingFuncs/parameterValueToIntSlice/stringSliceToIntSlice"
)

func ParameterValueToSortedIntSlice(parameterValue string) ([]int, error) {

	stringSlice, err := converting.ParameterValueToStringSlice(parameterValue)
	if err != nil {
		return nil, err
	}

	intSlice, err := conv.StringSliceToIntSlice(parameterValue, stringSlice)
	if err != nil {
		return nil, err
	}

	sort.Ints(intSlice)
	return intSlice, nil
}
