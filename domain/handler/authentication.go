package handler

import (
	"github.com/gofiber/fiber/v2"
)

type AuthenticationHandler struct{}

func (auth *AuthenticationHandler) Login(c *fiber.Ctx) error {
	return c.JSON(nil)
}
