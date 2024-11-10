package convfuncs

import (
	"errors"
	"log"
	"strconv"
)

var ErrUnsupported = errors.New("unsupported value")

func stringToInteger(parameterValue string) int {
	intParameterValue, err := strconv.Atoi(string(parameterValue))
	if err != nil {
		log.Printf("%s: %s", ErrUnsupported, parameterValue)
		return 0
	}
	return intParameterValue
}
