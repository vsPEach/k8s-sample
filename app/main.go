package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"status": "ok",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
