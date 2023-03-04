package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func MerchantDashboard(c *fiber.Ctx) error {
	return c.Render("merchant-dashboard", fiber.Map{
		"Username": "Sausagenoods",
		"CurrentBalance": "12",
	}, "layouts/merchant-panel")
}
