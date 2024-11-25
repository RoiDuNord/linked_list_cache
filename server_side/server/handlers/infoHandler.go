package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/cache"
	"server/config"
	"server/db"
	"server/server/checkers"
	converting "server/server/convertingResponseToIntSlice"
)

type Server struct {
	db    *db.Database
	cache *cache.Cache

	curFactor int
}

type UserResponse struct {
	Data []int `json:"userResponse"`
}

func New(cfg config.Config, db *db.Database, cache *cache.Cache) *Server {
	return &Server{
		db:        db,
		cache:     cache,
		curFactor: cfg.Factor,
	}
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
		log.Printf("requested values %d fully cached\n", inputNumbers)
	}

	dbData := s.uncachedDataHandler(uncachedNumbers)
	responseData := append(cacheData, dbData...)
	userResponseHandler(responseData, w)
}

func responseValue(parameter string, r *http.Request) converting.GetInfoRequestData {
	return converting.GetInfoRequestData(r.URL.Query().Get(parameter))
}

func userResponseHandler(responseData []int, w http.ResponseWriter) {
	userResponse := UserResponse{
		Data: responseData,
	}
	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
}

func (s *Server) uncachedDataHandler(uncachedNumbers []int) []int {
	dbData := s.db.FindNumbers(uncachedNumbers, s.curFactor)

	for i := range uncachedNumbers {
		s.cache.InsertNumber(uncachedNumbers[i], dbData[i])
	}

	return dbData
}
