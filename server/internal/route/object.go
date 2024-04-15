package route

import (
	"server/internal/app"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func ObjectRoute(handler *app.ServiceServer, enforcer *casbin.Enforcer) *fiber.App {
	router := fiber.New()
	// route.Get("/", handler.ListExample)
	router.Post("/", handler.CreateObject)
	router.Post("/http", adaptor.HTTPHandlerFunc(handler.CreateObjectHttp))
	return router
}
