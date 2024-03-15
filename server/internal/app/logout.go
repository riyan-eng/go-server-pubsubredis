package app

import (
	"server/infrastructure"
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Logout
// @Tags       	Authentication
// @Produce		json
// @Router		/auth/logout/ [delete]
// @Security ApiKeyAuth
func (s *ServiceServer) Logout(c *fiber.Ctx) error {
	userUUID := c.Locals("user_uuid").(string)
	s.authService.Logout(entity.AuthenticationLogoutReq{
		UserUUID: userUUID,
	})
	return util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_CREATE"))
}
