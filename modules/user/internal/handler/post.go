package handler

import (
	"simple-modular/internal/system"
	"simple-modular/modules/user/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h handler) CreateNewUser(c *fiber.Ctx) error {
	var req domain.UserRequest

	if err := c.BodyParser(&req); err != nil {
		var resp = system.ErrorBadRequest()
		resp.SetError(err)

		return c.Status(resp.StatusCode).JSON(resp)
	}

	var user = req.ToDomain()
	if err := h.svc.Create(c.UserContext(), user); err != nil {
		var resp = system.ErrorInternal()
		resp.SetError(err)

		return c.Status(resp.StatusCode).JSON(resp)
	}
	var resp = system.SuccessCreated()
	return c.Status(resp.StatusCode).JSON(resp)
}
