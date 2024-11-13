package handlers

import (
	"workWithCache/cache"
	"workWithCache/config"
	"workWithCache/db"
)

type Server struct {
	db    *db.Database
	cache *cache.Cache

	currFactor int
}

func New(cfg config.Config, db *db.Database, cache *cache.Cache) *Server {
	return &Server{
		db:         db,
		cache:      cache,
		currFactor: cfg.Factor,
	}
}
