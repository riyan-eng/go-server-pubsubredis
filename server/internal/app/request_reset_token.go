package app

import (
	"server/internal/dto"
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Request Reset Token
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthenticationRequestResetToken	true  "body"
// @Router		/auth/request_reset_token/ [post]
func (s *ServiceServer) RequestResetToken(c *fiber.Ctx) error {
	body := new(dto.AuthenticationRequestResetToken)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	s.authService.RequestResetToken(entity.AuthenticationRequestResetToken{
		Email:  body.Email,
		Issuer: string(c.Request().Host()),
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_REQUEST_TOKEN_RESET)
}
