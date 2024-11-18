package valuetostringslice

import (
	"strings"
	"workWithCache/server/convertingResponseToIntSlice/convertingFuncs/parameterValueToIntSlice/parameterValueToStringSlice/deduplication"
	customerrors "workWithCache/server/customErrors"
)

func ParameterValueToStringSlice(parameter string) ([]string, error) {
	cleanedParam := strings.Trim(parameter, "[]{}")
	trimmedParam := strings.TrimSpace(cleanedParam)
	stringSlice := strings.Split(trimmedParam, ",")
	if len(stringSlice) > 10 {
		return nil, customerrors.ExceedingTenValuesError(stringSlice)
	}
	deduplicatedSlice := deduplication.DeduplicateStringSlice(stringSlice)
	return deduplicatedSlice, nil
}
