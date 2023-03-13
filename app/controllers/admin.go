package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/moneropay/go-monero/walletrpc"

	"gitlab.com/moneropay/metronero/metronero-frontend/app/api"
)

func AdminDashboard(c *fiber.Ctx) error {
	token := c.Cookies("token")
	resp, err := api.GetAdminDashboard(token)
	if err != nil {
		return serveErrorPage(c, http.StatusInternalServerError, err.Error())
	}
	return c.Render("admin-dashboard", fiber.Map{
		"PageTitle":   "Dashboard",
		"Balance":     walletrpc.XMRToDecimal(resp.Stats.WalletBalance),
		"Profits":     walletrpc.XMRToDecimal(resp.Stats.TotalProfits),
		"Withdrawals": resp.Recent,
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
		"Username":  params["uname"],
		"PageTitle": "Merchant Edit",
	}, "layouts/admin-panel")
}
