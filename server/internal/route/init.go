package route

import (
	"server/config"
	"server/internal/app"
	"server/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupSubApp(fiberApp *fiber.App,
	exampleService service.ExampleService,
	authenticationService service.AuthenticationService,
	objectService service.ObjectService,
) {
	allHandler := app.NewService(exampleService, authenticationService, objectService)
	enforcer := config.NewEnforcer()

	// mounting sub app
	fiberApp.Mount("/", Index(allHandler, enforcer))
	fiberApp.Mount("/example", Example(allHandler, enforcer))
	fiberApp.Mount("/auth", Authentication(allHandler, enforcer))
	fiberApp.Mount("/template", Template(allHandler, enforcer))
	fiberApp.Mount("/import", Import(allHandler, enforcer))
	fiberApp.Mount("/export", Export(allHandler, enforcer))
	fiberApp.Mount("/export", Export(allHandler, enforcer))
	fiberApp.Mount("/object", ObjectRoute(allHandler, enforcer))
}
