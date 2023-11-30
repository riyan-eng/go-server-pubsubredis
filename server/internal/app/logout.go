package app

import (
	dtoservice "server/internal/dto_service"
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
	s.authService.Logout(dtoservice.AuthenticationLogoutReq{
		UserUUID: userUUID,
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_LOGOUT)
}
