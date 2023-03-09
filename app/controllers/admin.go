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
