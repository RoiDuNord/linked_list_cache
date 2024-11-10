package handlers

import (
	"fmt"
	"net/http"
)

func (s *Server) OutputHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Third step\n")
	s.cache.Data.Output(w)
}
