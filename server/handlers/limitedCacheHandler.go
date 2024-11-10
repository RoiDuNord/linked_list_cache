package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) LimitedCacheHandler(w http.ResponseWriter, r *http.Request) {
	newCacheSize := -7

	if newCacheSize == 0 {
		fmt.Fprint(w, "Cache size set to zero, but that's okay.\n")
		return
	}

	err := s.cache.LimitingNodesQuantity(newCacheSize)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Second step\n")
	s.cache.Data.Output(w)
}
