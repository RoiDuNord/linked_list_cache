package server

import (
	"fmt"
	"log"
	"net/http"
	"server/cache"
	"server/config"
	"server/db"
	"server/server/handlers"

	"github.com/go-chi/chi"
)

func Run(cfg config.Config) error {
	router := chi.NewRouter()

	database := db.New(cfg)
	cache := cache.New(cfg)

	s := handlers.New(cfg, database, cache)

	go s.StartFactorizeIncrement()

	router.Get("/info", s.InfoHandler)
	router.Post("/admin/cache/changeSize", s.ChangeCacheSize)
	router.Post("/admin/cache/changeFactor", s.ChangeFactor)
	router.Get("/admin/cache/output", s.OutputHandler)
	router.Post("/admin/worker/setActiveStatus", s.StartStopWorker)

	log.Printf("Starting HTTP server on port %d", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}

	return nil
}
