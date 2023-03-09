package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

func GetRegister(c *fiber.Ctx) error {
	return c.Render("register", nil)
}
