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

	db := db.New(cfg)

	cache := cache.New(cfg)

	s := handlers.New(cfg, db, cache)

	go s.StartFactorizeIncrement()

	router.Get("/info", s.InfoHandler)
	router.Post("/admin/cache/changeSize", s.ChangeCacheSize)
	router.Post("/admin/cache/changeFactor", s.ChangeFactor)
	router.Get("/admin/cache/output", s.OutputHandler)
	router.Post("/admin/worker/setActiveStatus", s.StartStopWorker)

	log.Printf("server runs at port %d", cfg.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		return err
	}
	return nil
}
