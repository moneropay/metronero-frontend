package main

import (
	"gitlab.com/moneropay/metronero/metronero-frontend/app/controllers"
	"gitlab.com/moneropay/metronero/metronero-frontend/utils/server"
	"gitlab.com/moneropay/metronero/metronero-frontend/utils/config"
)

func main() {
	cfg := config.Load()
	app := server.Init(cfg)

	app.Get("/login", controllers.GetLogin)
	app.Post("/login", controllers.PostLogin)

	merchant := app.Group("/merchant")
	merchant.Get("/dashboard", controllers.MerchantDashboard)

	app.Static("/assets", "./public")

	app.StartServerWithGracefulShutdown()
}
