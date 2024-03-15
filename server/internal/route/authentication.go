package route

import (
	"server/internal/app"
	"server/internal/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func Authentication(handler *app.ServiceServer, enforcer *casbin.Enforcer) *fiber.App {
	router := fiber.New()
	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
	router.Post("/refresh_token", handler.RefreshToken)
	router.Post("/request_reset_token", handler.RequestResetToken)
	router.Post("/validate_reset_token", handler.ValidateResetToken)
	router.Post("/reset_password", handler.ResetPassword)

	// middleware auth
	router.Use(middleware.AuthorizeJwt())

	router.Delete("/logout", handler.Logout)
	router.Get("/me", handler.Me)
	return router
}
