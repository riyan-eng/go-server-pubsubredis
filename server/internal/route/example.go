package route

import (
	"server/internal/app"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func Example(handler *app.ServiceServer, enforcer *casbin.Enforcer) *fiber.App {
	router := fiber.New()
	router.Get("/", handler.ListExample)
	router.Post("/", handler.CreateExample)
	router.Get("/:id", handler.DetailExample)
	router.Put("/:id", handler.PutExample)
	router.Patch("/:id", handler.PatchExample)
	router.Delete("/:id", handler.DeleteExample)
	return router
}
