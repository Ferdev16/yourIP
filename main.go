package main

import (
	"github.com/Ferdev16/yourIP/ip"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		res, err := ip.GetIP()
		if err != nil {
			return fiber.NewError(405, "Tu ip no fue encontrada")
		}
		return c.SendString(res)
	})

	app.Listen(":8080")
}
