package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/healt", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"status": "ok",
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"sosi": "huy",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
