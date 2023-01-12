package middleware

import (
	"errors"
	"fmt"
	"simple-api/domain/dto"
	"simple-api/helpers/jwt"
	"simple-api/helpers/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationFilter(c *fiber.Ctx) error {
	authorization_header := c.Get("Authorization")
	if authorization_header == "" {
		return response.FailOnError(c, errors.New("AUTHORIZATION IS MISSING"), fiber.ErrUnauthorized.Code)
	}
	token := strings.Split(authorization_header, "Bearer ")[1]
	user_info, err := jwt.Decode(token)
	if err != nil {
		return response.FailOnError(c, err, fiber.ErrUnauthorized.Code)
	}

	dto.UserId = fmt.Sprintf("%s", user_info["id"])
	dto.AirlineId = fmt.Sprintf("%s", user_info["airline_id"])

	return c.Next()
}
