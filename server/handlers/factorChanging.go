package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	customerrors "workWithCache/server/customErrors"
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
	// json.NewEncoder(w).Encode(ChangedFactor(factor))
	log.Printf("new factor: %d", factor.Value)

	s.curFactor = factor.Value
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

func (s *Server) StartStopWorker(w http.ResponseWriter, r *http.Request) {
	isActive = !isActive
	if isActive {
		w.WriteHeader(http.StatusOK)
		log.Println("worker started")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("worker stopped")
	}
}
