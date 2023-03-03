package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"gitlab.com/moneropay/metronero-frontend/app/config"
)

func main() {
	engine := html.New("./views", ".html")

	// Reload the templates on each render, good for development
	//engine.Reload(true) // Optional. Default: false

	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	// To render a template, you can call the ctx.Render function
	// Render(tmpl string, values interface{}, layout ...string)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("merchant-dashboard", fiber.Map{
			"Username": "Sausagenoods",
			"CurrentBalance": "12",
		}, "layouts/merchant-panel")
	})

	log.Fatal(app.Listen(config.Bind))
}
