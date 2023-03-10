package main

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"

	"gitlab.com/moneropay/metronero/metronero-frontend/app/controllers"
	"gitlab.com/moneropay/metronero/metronero-frontend/utils/config"
	"gitlab.com/moneropay/metronero/metronero-frontend/utils/server"
)

func main() {
	config.Load()
	app := server.Init()

	app.Static("/assets", "./public")

	app.Get("/login", controllers.GetLogin)
	app.Post("/login", controllers.PostLogin)
	app.Get("/register", controllers.GetRegister)
	app.Post("/register", controllers.PostRegister)
	app.Get("/logout", controllers.GetLogout)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  []byte(config.JwtSecret),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect("/login?expired=true")
		},
	}))

	merchant := app.Group("/merchant")
	merchant.Use(server.MerchantJwtMiddleware)
	merchant.Get("/auth", controllers.MerchantAuth)
	merchant.Get("/dashboard", controllers.MerchantDashboard)
	merchant.Get("/payments", controllers.MerchantPayments)
	merchant.Get("/withdrawals", controllers.MerchantWithdrawals)
	merchant.Get("/theme", controllers.MerchantTheme)

	admin := app.Group("/admin")
	admin.Use(server.AdminJwtMiddleware)
	admin.Get("/dashboard", controllers.AdminDashboard)
	admin.Get("/instance", controllers.AdminInstance)
	admin.Get("/withdrawals", controllers.AdminWithdrawals)
	admin.Get("/payments", controllers.AdminPayments)
	admin.Get("/merchants", controllers.AdminMerchants)
	admin.Get("/merchants/edit/:id", controllers.AdminMerchantEdit)

	app.StartServerWithGracefulShutdown()
}
