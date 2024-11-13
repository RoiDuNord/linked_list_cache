package checkers

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusBadRequest)
		e := userError{
			ErrorText: fmt.Sprintf("this URL doesn't contain requested parameter '%s'", parameter),
		}
		bytes, _ := json.Marshal(e)
		w.Write(bytes)
		return false
	}
	return true
}

type userError struct {
	ErrorText string `json:"error"`
}
