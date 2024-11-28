package handlers

import (
	"encoding/json"
	"net/http"
	customerrors "server/server/customErrors"
)

type NewCacheSize struct {
	Size int `json:"newSize"`
}

type ChangedSize struct {
	Size int `json:"changedSize"`
}

func (s *Server) ChangeCacheSize(w http.ResponseWriter, r *http.Request) {
	var newSizeOfCache NewCacheSize
	if err := json.NewDecoder(r.Body).Decode(&newSizeOfCache); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	err := s.cache.LimitingNodesQuantity(newSizeOfCache.Size)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := customerrors.WrongSizeError(err)
		json.NewEncoder(w).Encode(e)
		return
	}

	json.NewEncoder(w).Encode(ChangedSize(newSizeOfCache))
	// log.Printf("new cache size: %d", newSizeOfCache.Size)
}
