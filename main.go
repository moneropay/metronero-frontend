package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"gitlab.com/moneropay/metronero-frontend/app/config"
	"gitlab.com/moneropay/metronero-frontend/app/controllers"
	"gitlab.com/moneropay/metronero-frontend/app/utils"
)

func main() {
	engine := html.New("./views", ".html")

	if config.Debug {
		engine.Reload(true)
		engine.Debug(true)
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	merchant := app.Group("/merchant")
	merchant.Get("/dashboard", controllers.MerchantDashboard)

	utils.StartServerWithGracefulShutdown(app, config.Bind)
}
