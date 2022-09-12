package main

import (
	"api/database"
	"api/middleware"
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	middleware.AppMiddleware(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	routes.RegisterRoutes(app)
	app.Listen(":8000")
}
