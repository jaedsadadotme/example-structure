package response

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FailOnError(c *fiber.Ctx, err error, status int, field ...[]interface{}) error {
	var errorFields []interface{}
	if len(field) > 0 {
		errorFields = field[0]
	} else {
		errorFields = nil
	}
	return c.Status(status).JSON(fiber.Map{
		"error":            http.StatusText(status),
		"errorStatus":      status,
		"errorDescription": err.Error(),
		"errorAt":          time.Now().String(),
		"errorTraceId":     c.GetRespHeader("X-Request-Id"),
		"errorUri":         "",
		"errorOn":          "",
		"errorFields":      errorFields,
		"errorData":        "",
		"state":            nil,
	})
}
