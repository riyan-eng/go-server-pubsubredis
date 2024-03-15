package app

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Reset Password
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthenticationResetPassword	true  "body"
// @Router		/auth/reset_password/ [post]
func (s *ServiceServer) ResetPassword(c *fiber.Ctx) error {
	body := new(dto.AuthenticationResetPassword)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("OK_CREATE"), fiber.StatusBadRequest)
	}

	s.authService.ResetPassword(entity.AuthenticationResetPasswordReq{
		ResetToken: body.ResetToken,
		Password:   body.Password,
	})
	return util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_CREATE"))
}
