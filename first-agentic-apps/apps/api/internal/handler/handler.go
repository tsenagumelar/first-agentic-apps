package handler

import (
	"github.com/gofiber/fiber/v2"

	"mvp-web-admin-api/internal/service"
)

type Handler struct {
	auth *service.AuthService
}

func NewHandler() *Handler {
	return &Handler{auth: service.NewAuthService()}
}

func (h *Handler) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "login endpoint ready"})
}

func (h *Handler) Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "register endpoint ready"})
}
