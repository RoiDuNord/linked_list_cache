package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) SetFactor(w http.ResponseWriter, r *http.Request) {
	// парсим пришедший json
	s.curFactor = 0
}

func (s *Server) StartStopWorker(w http.ResponseWriter, r *http.Request) {
	// парсим пришедший json
	// вкл/выкл горутина
}

func (s *Server) StartFactorizeIncrement() {
	for {
		time.Sleep(5 * time.Second)
		s.curFactor++
		fmt.Printf("Множитель обновлён: %d\n", s.curFactor)
		cur := s.cache.List.Head
		for cur != nil {
			cur.Data.MultipliedNumber = cur.Data.Number * s.curFactor
			cur = cur.Next
		}
	}
}
