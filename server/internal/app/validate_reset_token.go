package app

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Validate Token Reset Password
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthenticationValidateResetToken	true  "body"
// @Router		/auth/validate_reset_token/ [post]
func (s *ServiceServer) ValidateResetToken(c *fiber.Ctx) error {
	body := new(dto.AuthenticationValidateResetToken)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("OK_CREATE"), fiber.StatusBadRequest)
	}

	s.authService.ValidateResetToken(entity.AuthenticationValidateResetTokenReq{
		ResetToken: body.ResetToken,
	})
	return util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_CREATE"))
}
