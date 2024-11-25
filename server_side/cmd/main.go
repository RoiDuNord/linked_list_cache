package main

import (
	"log"
	"server/config"
	"server/server"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	if cfg == (config.Config{}) {
		log.Fatal("Empty config")
	}

	server.Run(cfg)
}
