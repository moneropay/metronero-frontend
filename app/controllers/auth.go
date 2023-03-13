package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"gitlab.com/moneropay/metronero/metronero-frontend/app/api"
)

func GetLogin(c *fiber.Ctx) error {
	// Display a success message if the user ended up here
	// after successfully logging in.
	reg := c.Query("success", "false")

	// Display a message telling the user to login to proceed
	// in case of invalid or expired token.
	exp := c.Query("expired", "false")
	return c.Render("login", fiber.Map{
		"Registered": reg,
		"Expired": exp,
	})
}

func PostLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return c.SendString("Bad request!")
	}
	token, err := api.UserLogin(username, password)
	if err != nil {
		return c.SendString(err.Error())
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token.Token
	t, _ := time.Parse(time.RFC3339, token.ValidUntil)
	cookie.Expires = t
	cookie.Secure = true
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	if username == "admin" {
		return c.Redirect("/admin/dashboard")
	}
	return c.Redirect("/merchant/dashboard")
}

func GetRegister(c *fiber.Ctx) error {
	return c.Render("register", nil)
}

func PostRegister(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return c.SendString("Bad request!")
	}
	if err := api.UserRegister(username, password); err != nil {
		return c.SendString(err.Error())
	}
	return c.Redirect("/login?success=true")
}

func GetLogout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.Redirect("/login")
}
