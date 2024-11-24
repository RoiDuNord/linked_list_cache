package checkers

import (
	"encoding/json"
	"net/http"
	converting "server/server/convertingResponseToIntSlice"
	customerrors "server/server/customErrors"
)

func InputNumbers(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := customerrors.ParameterIsNotNumbersOrEmptyValueError(err)
		json.NewEncoder(w).Encode(e)
		return false
	}
	return true
}

func ParameterValue(parameterValue converting.GetInfoRequestData, parameter string, w http.ResponseWriter) bool {
	if parameterValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		e := customerrors.URLWithoutParameterNumbersError(parameter)
		json.NewEncoder(w).Encode(e)
		return false
	}
	return true
}

type userError struct {
	ErrorText string `json:"error"`
}
