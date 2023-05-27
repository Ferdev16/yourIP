package main

import (
	"net/http"

	"github.com/Ferdev16/yourIP/ip"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("./template", ".html")

	app := fiber.New(fiber.Config{
		Views:   engine,
		Prefork: true,
	})

	app.Use("/template", filesystem.New(filesystem.Config{
		Root:   http.Dir("./template"),
		Browse: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"ipPublic":  ip.GetIP(),
			"ipPrivate": c.IP(),

			"path": "./template",
		})
	})

	app.Listen(":8080")
}
