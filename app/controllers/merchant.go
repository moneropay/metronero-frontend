package controllers

import (
	"github.com/gofiber/fiber/v2"

	"gitlab.com/moneropay/metronero/metronero-frontend/utils/token"
)

func MerchantDashboard(c *fiber.Ctx) error {
	return c.Render("merchant-dashboard", fiber.Map{
		"Username":       token.GetUsername(c),
		"PageTitle":      "Dashboard",
		"CurrentBalance": "12",
	}, "layouts/merchant-panel")
}

func MerchantPayments(c *fiber.Ctx) error {
	return c.Render("merchant-payments", fiber.Map{
		"Username":  token.GetUsername(c),
		"PageTitle": "Payments",
	}, "layouts/merchant-panel")
}

func MerchantWithdrawals(c *fiber.Ctx) error {
	return c.Render("merchant-withdrawals", fiber.Map{
		"Username":  token.GetUsername(c),
		"PageTitle": "Withdrawals",
	}, "layouts/merchant-panel")
}

func MerchantTheme(c *fiber.Ctx) error {
	return c.Render("merchant-theme", fiber.Map{
		"Username":  token.GetUsername(c),
		"PageTitle": "Theme",
	}, "layouts/merchant-panel")
}

func MerchantAuth(c *fiber.Ctx) error {
	return c.Render("merchant-auth", fiber.Map{
		"Username":  token.GetUsername(c),
		"PageTitle": "Authentication",
	}, "layouts/merchant-panel")
}
