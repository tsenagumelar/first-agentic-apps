package main

import (
	"log"

	"mvp-web-admin-api/internal/app"
	"mvp-web-admin-api/internal/config"
)

func main() {
	cfg := config.Load()
	server := app.NewServer(cfg)

	if err := server.Listen(cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
