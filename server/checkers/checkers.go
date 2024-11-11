package checkers

import (
	"fmt"
	"log"
	"net/http"
	converting "workWithCache/server/convertingResponseToIntSlice"
)

func Database(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("%v", err)
		fmt.Fprintln(w, err)
		return false
	}
	return true
}

func InputNumbers(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("%v", err)
		fmt.Fprintln(w, err)
		return false
	}
	return true
}

func ParameterValue(parameterValue converting.GetInfoRequestData, parameter string, w http.ResponseWriter) bool {
	if parameterValue == "" {
		http.Error(w, fmt.Sprintf("this URL doesn't contain requested parameter '%s'", parameter), http.StatusBadRequest)
		return false
	}
	return true
}
