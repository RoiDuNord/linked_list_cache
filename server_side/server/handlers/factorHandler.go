package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	customerrors "server/server/customErrors"
	"time"
)

type Factor struct {
	Value int `json:"factor"`
}

type ChangedFactor struct {
	Value int `json:"changedFactor"`
}

func (s *Server) ChangeFactor(w http.ResponseWriter, r *http.Request) {
	var factor Factor
	if err := json.NewDecoder(r.Body).Decode(&factor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	if factor.Value < 0 {
		w.WriteHeader(http.StatusBadRequest)
		e := customerrors.NegativeFactorError(factor.Value)
		json.NewEncoder(w).Encode(customerrors.WrongFactorError(e))
		return
	}
	json.NewEncoder(w).Encode(ChangedFactor(factor))

	s.curFactor = factor.Value
	log.Printf("new factor %d", s.curFactor)
}

var isActive bool

func (s *Server) StartFactorizeIncrement() {
	isActive = true

	for {
		if !isActive {
			continue
		}
		time.Sleep(5 * time.Second)
		s.curFactor++
		log.Printf("incremented factor: %d\n", s.curFactor)
		s.cache.UpdateWithNewFactor(s.curFactor)
	}
}

type ActiveWorker struct {
	Activity bool `json:"isActiveWorker"`
}

func (s *Server) StartStopWorker(w http.ResponseWriter, r *http.Request) {
	isActive = !isActive
	if isActive {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(ActiveWorker{Activity: isActive})
}
