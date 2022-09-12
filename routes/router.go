package routes

import (
	"api/routes/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func RegisterRoutes(app *fiber.App) {
	_auth := app.Group("/auth")
	auth.AuthRouter(_auth)

	app.Get("/dashboard", monitor.New())
}
