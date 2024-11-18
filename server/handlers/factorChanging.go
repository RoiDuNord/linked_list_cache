package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) FactorChanging(w http.ResponseWriter, r *http.Request) {
	go s.factorModifiedList()
}

func (s *Server) factorModifiedList() {
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
