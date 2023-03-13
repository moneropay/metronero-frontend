package controllers

import "github.com/gofiber/fiber/v2"

func serveErrorPage(c *fiber.Ctx, code int, message string) error {
	return c.Render("error", fiber.Map{
		"Code": code,
		"Message": message,
	})
}
