package main

import (
	"gitlab.com/moneropay/metronero-frontend/app/controllers"
	"gitlab.com/moneropay/metronero-frontend/utils/server"
	"gitlab.com/moneropay/metronero-frontend/utils/config"
)

func main() {
	cfg := config.Load()
	app := server.Init(cfg)

	merchant := app.Group("/merchant")
	merchant.Get("/dashboard", controllers.MerchantDashboard)

	app.Static("/assets", "./public")

	app.StartServerWithGracefulShutdown()
}
