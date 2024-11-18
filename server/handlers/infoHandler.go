package handlers

import (
	"encoding/json"
	"net/http"
	"workWithCache/cache"
	"workWithCache/server/checkers"
	converting "workWithCache/server/convertingResponseToIntSlice"
	customerrors "workWithCache/server/customErrors"
)

func responseValue(parameter string, r *http.Request) converting.GetInfoRequestData {
	return converting.GetInfoRequestData(r.URL.Query().Get(parameter))
}

func (s *Server) InfoHandler(w http.ResponseWriter, r *http.Request) {
	parameter := "numbers"
	parameterValue := responseValue(parameter, r)
	if !checkers.ParameterValue(parameterValue, parameter, w) {
		return
	}

	inputNumbers, err := parameterValue.ToIntSlice(parameter)
	if !checkers.InputNumbers(err, w) {
		return
	}

	cacheData, uncachedNumbers := s.cache.FindNumbers(inputNumbers)
	if len(uncachedNumbers) == 0 {
		e := cache.ZeroLenUncached(inputNumbers)
		json.NewEncoder(w).Encode(e)
		userResponseHandler(cacheData, w)
		return
	} else {
		dbData := s.uncachedDataHandler(uncachedNumbers)
		responseData := append(cacheData, dbData...)
		userResponseHandler(responseData, w)
	}

	// s.cache.List.Output(w)
}

type UserResponse struct {
	Response []int `json:"userResponse"`
}

func userResponseHandler(responseData []int, w http.ResponseWriter) {
	userResponse := UserResponse{
		Response: responseData,
	}
	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e := customerrors.JSONEncodingError(err)
		json.NewEncoder(w).Encode(e)
	}
}

func (s *Server) uncachedDataHandler(uncachedNumbers []int) []int {
	dbData := s.db.FindNumbers(uncachedNumbers, s.curFactor)

	for i := range uncachedNumbers {
		s.cache.InsertNumber(uncachedNumbers[i], dbData[i])
	}

	return dbData
}
