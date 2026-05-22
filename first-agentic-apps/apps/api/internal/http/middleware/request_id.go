package middleware

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/fiber/v2/middleware/requestid"

func RequestID() fiber.Handler {
	return requestid.New()
}
