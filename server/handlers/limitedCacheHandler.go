package handlers

import (
	"encoding/json"
	"net/http"
	customerrors "workWithCache/server/customErrors"
)

type NewCacheSize struct {
	Size int `json:"newSize"`
}

type ChangedSize struct {
	Size int `json:"changedSize"`
}

func (s *Server) LimitedCacheHandler(w http.ResponseWriter, r *http.Request) {
	var newSizeOfCache NewCacheSize

	err := json.NewDecoder(r.Body).Decode(&newSizeOfCache)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := customerrors.JSONEncodingError(err)
		json.NewEncoder(w).Encode(e)
		return
	}
	defer r.Body.Close()

	err = s.cache.LimitingNodesQuantity(newSizeOfCache.Size)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := customerrors.WrongSizeError(err)
		json.NewEncoder(w).Encode(e)
		return
	}

	json.NewEncoder(w).Encode(ChangedSize(newSizeOfCache))
}
