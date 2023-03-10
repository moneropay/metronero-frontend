package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func MerchantDashboard(c *fiber.Ctx) error {
	return c.Render("merchant-dashboard", fiber.Map{
		"Username": "Sausagenoods",
		"PageTitle": "Dashboard",
		"CurrentBalance": "12",
	}, "layouts/merchant-panel")
}

func MerchantPayments(c *fiber.Ctx) error {
	return c.Render("merchant-payments", fiber.Map{
		"Username": "Sausagenoods",
		"PageTitle": "Payments",
	}, "layouts/merchant-panel")
}

func MerchantWithdrawals(c *fiber.Ctx) error {
	return c.Render("merchant-withdrawals", fiber.Map{
		"Username": "Sausagenoods",
		"PageTitle": "Withdrawals",
	}, "layouts/merchant-panel")
}

func MerchantTheme(c *fiber.Ctx) error {
	return c.Render("merchant-theme", fiber.Map{
		"Username": "Sausagenoods",
		"PageTitle": "Theme",
	}, "layouts/merchant-panel")
}

func MerchantAuth(c *fiber.Ctx) error {
	return c.Render("merchant-auth", fiber.Map{
		"Username": "Sausagenoods",
		"PageTitle": "Authentication",
	}, "layouts/merchant-panel")
}
