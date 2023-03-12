package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func MerchantJwtMiddleware(c *fiber.Ctx) error {
	cookie := c.Locals("user")
	if cookie == nil {
		return c.Redirect("/login")
	}
	token := cookie.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	if claims["username"].(string) == "admin" {
		return c.Redirect("/login")
	}
	return c.Next()
}

func AdminJwtMiddleware(c *fiber.Ctx) error {
	cookie := c.Locals("user")
	if cookie == nil {
		return c.Redirect("/login")
	}
	token := cookie.(*jwt.Token)
	//token := c.Locals("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	if claims["username"].(string) != "admin" {
		return c.Redirect("/login")
	}
	return c.Next()
}
