package route

import (
	"server/internal/app"
	"server/internal/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func ExampleRoute(a *fiber.App, handler *app.ServiceServer, enforcer *casbin.Enforcer) {
	a.Get("/example-template", handler.TemplateExample)
	a.Get("/example-download", handler.DownloadExample)

	route := a.Group("/example", middleware.AuthorizeJwt(), middleware.PermitCasbin(enforcer))
	route.Get("/", handler.ListExample)
	route.Post("/", handler.CreateExample)
	route.Get("/:id", handler.DetailExample)
	route.Put("/:id", handler.PutExample)
	route.Patch("/:id", handler.PatchExample)
	route.Delete("/:id", handler.DeleteExample)
	route.Delete("/:id", handler.DeleteExample)

}
