package handlers

import (
	"net/http"
)

func (s *Server) OutputHandler(w http.ResponseWriter, r *http.Request) {
	s.cache.List.Output(w)
}
