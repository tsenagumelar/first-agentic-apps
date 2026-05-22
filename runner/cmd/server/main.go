package main

import (
	"log"

	"first-agentic-runner/internal/config"
	"first-agentic-runner/internal/handler"
	"first-agentic-runner/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		AppName: "first-agentic-codex-runner",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	codexService := service.NewCodexService(cfg.CodexBinary, cfg.Timeout)
	codexHandler := handler.NewCodexHandler(codexService)

	app.Get("/health", codexHandler.Health)
	app.Post("/run-codex", codexHandler.RunCodex)

	log.Println("Codex runner running on http://localhost:" + cfg.AppPort)
	log.Println("Codex binary:", cfg.CodexBinary)

	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
