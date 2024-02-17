package app

import (
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Me
// @Tags       	Authentication
// @Produce		json
// @Router		/auth/me/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) Me(c *fiber.Ctx) error {
	userUUID := c.Locals("user_uuid").(string)
	service := s.authService.Me(entity.AuthenticationMeReq{
		UserUUID: userUUID,
	})
	return util.NewResponse(c).Success(service.Data, nil, util.MESSAGE_OK_READ)
}
