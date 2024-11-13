package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"workWithCache/config"
	"workWithCache/server/checkers"
	converting "workWithCache/server/convertingResponseToIntSlice"
	convert "workWithCache/server/convertingResponseToIntSlice/convertingFuncs/intSliceToStringSlice"
)

func responseValue(parameter string, _ http.ResponseWriter, r *http.Request) converting.GetInfoRequestData {
	return converting.GetInfoRequestData(r.URL.Query().Get(parameter))
}

func (s *Server) InfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "First step\n")

	parameter := "numbers"
	parameterValue := responseValue(parameter, w, r)
	if !checkers.ParameterValue(parameterValue, parameter, w) {
		return
	}

	inputNumbers, err := parameterValue.ToIntSlice(parameter)
	if !checkers.InputNumbers(err, w) {
		return
	}

	cfg, err := config.ParseConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	cache, uncachedNumbers := s.cache.FindNumbers(inputNumbers)
	fmt.Fprint(w, "Cache: ", cache, "\n")

	if len(uncachedNumbers) == 0 {
		fmt.Fprintf(w, "User Response: %s\n", strings.Join(convert.IntSliceToStringSlice(cache), ", "))
		http.Error(w, "No uncached numbers provided", http.StatusBadRequest)
		return
	}

	dbData, err := s.db.FindNumbers(inputNumbers, cfg.Factor)
	if !checkers.Database(err, w) {
		return
	}

	fmt.Fprint(w, "DB: ", dbData, "\n")

	for i := range uncachedNumbers {
		s.cache.InsertNumber(uncachedNumbers[i], dbData[i])
	}

	responseData := append(cache, dbData...)

	fmt.Fprintf(w, "User Response: %s\n", strings.Join(convert.IntSliceToStringSlice(responseData), ", "))

	s.cache.Data.Output(w)

}
