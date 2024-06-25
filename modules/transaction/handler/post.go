package handler

import (
	"simple-modular/internal/system"
	"simple-modular/modules/transaction/domain"

	"github.com/gofiber/fiber/v2"
)

func (h handler) Transfer(c *fiber.Ctx) error {
	var req = domain.TransferRequest{}

	if err := c.BodyParser(&req); err != nil {
		var resp = system.ErrorBadRequest()
		resp.SetError(err)

		return c.Status(resp.StatusCode).JSON(resp)
	}

	var trx = req.ToDomain()
	if err := h.svc.Transfer(c.UserContext(), trx); err != nil {
		var resp = system.ErrorInternal()
		resp.SetError(err)

		return c.Status(resp.StatusCode).JSON(resp)
	}

	var resp = system.SuccessCreated()
	return c.Status(resp.StatusCode).JSON(resp)
}

func (h handler) Topup(c *fiber.Ctx) error {
	var req = domain.TopupRequest{}

	if err := c.BodyParser(&req); err != nil {
		var resp = system.ErrorBadRequest()
		resp.SetError(err)

		return c.Status(resp.StatusCode).JSON(resp)
	}

	var trx = req.ToDomain()
	if err := h.svc.Topup(c.UserContext(), trx); err != nil {
		var resp = system.ErrorInternal()
		resp.SetError(err)

		return c.Status(resp.StatusCode).JSON(resp)
	}

	var resp = system.SuccessCreated()
	return c.Status(resp.StatusCode).JSON(resp)
}
