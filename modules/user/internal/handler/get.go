package handler

import (
	"net/http"
	"simple-modular/internal/system"
	"simple-modular/modules/user/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAll(c *fiber.Ctx) error {
	users, err := h.svc.GetAll()
	if err != nil {

		return c.Status(http.StatusInternalServerError).JSON(system.ErrorInternal())
	}

	var usersResponse []domain.UserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, domain.NewUserResponseFromDomain(user))
	}

	resp := system.Success()
	resp.SetPayload(usersResponse)

	return c.Status(resp.StatusCode).JSON(resp)
}
