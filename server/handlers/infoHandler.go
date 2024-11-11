package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"workWithCache/server/checkers"
	converting "workWithCache/server/convertingResponseToIntSlice"
	convert "workWithCache/server/convertingResponseToIntSlice/convertingFuncs/intSliceToStringSlice"
)

func responseValue(parameter string, _ http.ResponseWriter, r *http.Request) converting.GetInfoRequestData {
	return converting.GetInfoRequestData(r.URL.Query().Get(parameter))
}

func (s *Server) InfoHandler(w http.ResponseWriter, r *http.Request) {
	parameter := "numbers"
	parameterValue := responseValue(parameter, w, r)
	if !checkers.ParameterValue(parameterValue, parameter, w) {
		return
	}

	inputNumbers, err := parameterValue.ToIntSlice(parameter)
	if !checkers.InputNumbers(err, w) {
		return
	}

	factor := 4

	cache := s.cache
	userResponse, err := s.db.FindNumbers(inputNumbers, factor, cache)
	if !checkers.Database(err, w) {
		return
	}

	fmt.Fprintf(w, "User Response: %s\n", strings.Join(convert.IntSliceToStringSlice(userResponse), ","))
	fmt.Fprint(w, "First step\n")

	cache.Data.Output(w)
}

// пытаемся найти в кэше, потому что это быстрее
// cache.Find(inputNumbers)
// идём в бд и "ищем" эти числа, если надо
// собираем результат
// кладём в кэш результаты
// отдаёт пользователю
