package app

import (
	dtoservice "server/internal/dto_service"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Me
// @Tags       	Authentication
// @Produce		json
// @Router		/auth/me/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) Me(c *fiber.Ctx) error {
	idUser := util.StringNumToInt(c.Locals("user_id").(string))

	service := s.authService.Me(dtoservice.AuthenticationMeReq{
		IDUser: idUser,
	})
	return util.NewResponse(c).Success(service.Data, nil, util.MESSAGE_OK_READ)
}
