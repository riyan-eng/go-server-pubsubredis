package app

import (
	dtoservice "server/internal/dto_service"
	httprequest "server/pkg/http.request"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Reset Password
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.AuthenticationResetPassword	true  "body"
// @Router		/auth/reset_password/ [post]
func (s *ServiceServer) ResetPassword(c *fiber.Ctx) error {
	body := new(httprequest.AuthenticationResetPassword)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	s.authService.ResetPassword(dtoservice.AuthenticationResetPasswordReq{
		ResetToken: body.ResetToken,
		Password:   body.Password,
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_CHANGE_PASSWORD)
}
