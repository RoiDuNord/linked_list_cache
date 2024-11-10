package server

import (
	"log"
	"net/http"
	"workWithCache/cache"
	"workWithCache/db"
	"workWithCache/server/handlers"

	"github.com/go-chi/chi"
)

func Run() {
	router := chi.NewRouter()

	s := handlers.New(db.New(), cache.New())

	router.Get("/info", s.InfoHandler)
	router.Post("/info/setCache", s.LimitedCacheHandler)
	router.Get("/info/setCache/output", s.OutputHandler)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
