package handler

import (
	"simple-api/domain/repository"
	"simple-api/helpers/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repository repository.UserRepository
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	_, err := h.repository.FindAll()
	if err != nil {
		return response.
			FailOnError(c, err, 500)
	}
	return c.
		Status(fiber.StatusOK).
		JSON(map[string]string{
			"message": "hello user",
		})
}
