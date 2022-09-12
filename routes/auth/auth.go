package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router) {

	// dummy route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from /auth")
	})
}
