package server

import (
	"fmt"
	"log"
	"net/http"
	"workWithCache/cache"
	"workWithCache/config"
	"workWithCache/db"
	"workWithCache/server/handlers"

	"github.com/go-chi/chi"
)

func Run(cfg config.Config) error {
	router := chi.NewRouter()

	db := db.New(cfg)

	cache := cache.New(cfg)

	s := handlers.New(cfg, db, cache)

	router.Get("/info", s.InfoHandler)
	router.Post("/admin/cache/setSize", s.LimitedCacheHandler)
	router.Post("/admin/cache/changeFactor", s.FactorChanging)
	router.Get("/admin/cache/output", s.OutputHandler)

	log.Printf("Server runs at port %d", cfg.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		return err
	}
	return nil
}
