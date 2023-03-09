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
	app.Get("/register", controllers.GetRegister)
	//app.Post("/login", controllers.PostLogin)

	merchant := app.Group("/merchant")
	merchant.Get("/dashboard", controllers.MerchantDashboard)
	merchant.Get("/payments", controllers.MerchantPayments)
	merchant.Get("/theme", controllers.MerchantTheme)

	admin := app.Group("/admin")
	admin.Get("/dashboard", controllers.AdminDashboard)
	admin.Get("/instance", controllers.AdminInstance)
	admin.Get("/withdrawals", controllers.AdminWithdrawals)

	app.Static("/assets", "./public")

	app.StartServerWithGracefulShutdown()
}
