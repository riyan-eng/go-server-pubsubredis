package route

import (
	"server/internal/app"
	"server/internal/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func Export(handler *app.ServiceServer, enforcer *casbin.Enforcer) *fiber.App {
	router := fiber.New()
	// middleware
	router.Use(middleware.AuthorizeJwt())
	router.Use(middleware.PermitCasbin(enforcer))

	// route
	router.Get("/example", handler.DownloadExample)
	return router
}
