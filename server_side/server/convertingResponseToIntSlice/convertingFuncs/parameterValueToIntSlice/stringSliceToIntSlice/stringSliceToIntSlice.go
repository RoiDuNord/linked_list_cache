package convfuncs

import (
	"strings"
	customerrors "server/server/customErrors"
)

func StringSliceToIntSlice(parameterValue string, stringSlice []string) ([]int, error) {
	newSlice := make([]int, 0, len(stringSlice))
	for _, el := range stringSlice {
		el = strings.TrimSpace(el)
		number := stringToInteger(el)
		if number == 0 {
			continue
		}
		newSlice = append(newSlice, number)
	}
	if err := customerrors.EmptySliceError(newSlice, parameterValue); err != nil {
		return nil, err
	}
	return newSlice, nil
}
