package route

import (
	"server/internal/app"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func Index(handler *app.ServiceServer, enforcer *casbin.Enforcer) *fiber.App {
	router := fiber.New()
	// middleware
	// router.Use(middleware.AuthorizeJwt())
	// router.Use(middleware.PermitCasbin(enforcer))

	// route
	router.Get("/pdf", handler.DownloadPDF)
	return router
}
