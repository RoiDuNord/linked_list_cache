package handlers

import (
	"workWithCache/cache"
	"workWithCache/db"
)

type Server struct {
	db    *db.Database
	cache *cache.Cache
}

func New(db *db.Database, cache *cache.Cache) *Server {
	return &Server{
		db:    db,
		cache: cache,
	}
}
