package handler

import (
	"time"

	"first-agentic-runner/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CodexHandler struct {
	CodexService *service.CodexService
}

func NewCodexHandler(codexService *service.CodexService) *CodexHandler {
	return &CodexHandler{
		CodexService: codexService,
	}
}

func (h *CodexHandler) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "ok",
		"service":   "first-agentic-codex-runner",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func (h *CodexHandler) RunCodex(c *fiber.Ctx) error {
	var input service.RunCodexInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid request body: " + err.Error(),
		})
	}

	result := h.CodexService.Run(input)

	if !result.Success {
		return c.Status(fiber.StatusOK).JSON(result)
	}

	return c.JSON(result)
}
