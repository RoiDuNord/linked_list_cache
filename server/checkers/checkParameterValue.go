package checkers

import (
	"fmt"
	"net/http"
	converting "workWithCache/server/convertingResponseToIntSlice"
)

func ParameterValue(parameterValue converting.GetInfoRequestData, parameter string, w http.ResponseWriter) bool {
	if parameterValue == "" {
		http.Error(w, fmt.Sprintf("This URL doesn't contain requested parameter '%s'", parameter), http.StatusBadRequest)
		return false
	}
	return true
}
