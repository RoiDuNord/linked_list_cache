package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) LimitedCacheHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Second step\n")

	var newSize int

	fmt.Print("New size: ")
	fmt.Scan(&newSize)

	if newSize == 0 {
		fmt.Fprint(w, "Cache size set to zero, but that's okay.\n")
		return
	}

	err := s.cache.LimitingNodesQuantity(newSize)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Internal Server Error (%d)", http.StatusInternalServerError)
		return
	}

	s.cache.Data.Output(w)
}
