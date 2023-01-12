package routes

import (
	"net/http"
	"os"
	"simple-api/domain/handler"
	"simple-api/domain/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Config() *fiber.App {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "",
		AppName:       "Simple Api",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.
				Status(err.(*fiber.Error).Code).
				JSON(map[string]interface{}{
					"error":            http.StatusText(err.(*fiber.Error).Code),
					"errorStatus":      err.(*fiber.Error).Code,
					"errorDescription": err.Error(),
					"errorAt":          time.Now().String(),
					"errorTraceId":     c.GetRespHeader("X-Request-Id"),
					"errorUri":         "",
					"errorOn":          "",
					"errorFields":      "",
					"errorData":        "",
					"state":            nil,
				})
		},
	})

	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.Config{
		Next:         nil,
		Format:       "[${method}] | ${status} | ${time} | ${locals:requestid} | ${latency} | ${path} | ${ip} | ${error} | ${body} \n",
		TimeFormat:   "02/Jan/2006 15:04",
		TimeZone:     "Asia/Bangkok",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stderr,
	}))

	return app
}

func New() *fiber.App {
	app := Config()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]string{
			"message": "hello",
		})
	})
	auth := handler.AuthenticationHandler{}
	auth_router := app.Group("auth")
	auth_router.Post("/login", auth.Login)

	api := app.Group("api/", middleware.AuthenticationFilter)
	handler := handler.Handler{}
	api.Get("/user", handler.FindAll)

	return app
}
