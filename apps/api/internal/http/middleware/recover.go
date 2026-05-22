package middleware

import (
	"github.com/gofiber/fiber/v2"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func Recover() fiber.Handler {
	return fiberRecover.New()
}
