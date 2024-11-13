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

func Run(cfg config.Config) {
	router := chi.NewRouter()

	db := db.New(cfg)

	cache := cache.New(cfg)

	s := handlers.New(cfg, db, cache)

	router.Get("/info", s.InfoHandler)
	router.Post("/admin/cache/setSize", s.LimitedCacheHandler)
	router.Get("/admin/cache/output", s.OutputHandler)

	//  /admin/cache/setSize POST body: {size:5}
	router.Post("/admin/factor/set", nil) // Этап 2.5. Должна менять множитель на любой >0
	// Этап 3. Далее изменять множитель каждые 5 сек на 1 увеличиватьё

	log.Printf("Server runs at port %d", cfg.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
