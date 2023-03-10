package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func AdminDashboard(c *fiber.Ctx) error {
	return c.Render("admin-dashboard", fiber.Map{
		"PageTitle": "Dashboard",
	}, "layouts/admin-panel")
}

func AdminInstance(c *fiber.Ctx) error {
	return c.Render("admin-instance", fiber.Map{
		"PageTitle": "Instance Settings",
	}, "layouts/admin-panel")
}

func AdminWithdrawals(c *fiber.Ctx) error {
	return c.Render("admin-withdrawals", fiber.Map{
		"PageTitle": "Withdrawals",
	}, "layouts/admin-panel")
}

func AdminPayments(c *fiber.Ctx) error {
	return c.Render("admin-payments", fiber.Map{
		"PageTitle": "Payments",
	}, "layouts/admin-panel")
}

func AdminMerchants(c *fiber.Ctx) error {
	return c.Render("admin-merchants", fiber.Map{
		"PageTitle": "Merchants",
	}, "layouts/admin-panel")
}

func AdminMerchantEdit(c *fiber.Ctx) error {
	params := c.AllParams()
	return c.Render("admin-merchant-edit", fiber.Map{
		"Username": params["uname"],
		"PageTitle": "Merchant Edit",
	}, "layouts/admin-panel")
}
