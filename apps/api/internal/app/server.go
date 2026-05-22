package app

import (
	"github.com/gofiber/fiber/v2"

	"mvp-web-admin-api/internal/config"
	"mvp-web-admin-api/internal/handler"
	mw "mvp-web-admin-api/internal/http/middleware"
)

func NewServer(cfg config.Config) *fiber.App {
	app := fiber.New()
	app.Use(mw.RequestID())
	app.Use(mw.Recover())

	h := handler.NewHandler()

	app.Get("/health", h.Health)
	api := app.Group("/api")
	api.Post("/auth/login", h.Login)
	api.Post("/auth/register", h.Register)

	return app
}
