package route

import (
	"server/config"
	"server/internal/app"
	"server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func NewRoute(fiberApp *fiber.App,
	exampleService service.ExampleService,
	authenticationService service.AuthenticationService,
	objectService service.ObjectService,
) {
	allHandler := app.NewService(exampleService, authenticationService, objectService)
	enforcer := config.NewEnforcer()
	ExampleRoute(fiberApp, allHandler, enforcer)
	AuthenticationRoute(fiberApp, allHandler, enforcer)
	ObjectRoute(fiberApp, allHandler, enforcer)
}
