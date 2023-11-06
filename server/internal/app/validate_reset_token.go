package app

import (
	dtoservice "server/internal/dto_service"
	httprequest "server/pkg/http.request"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Validate Token Reset Password
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.AuthenticationValidateResetToken	true  "body"
// @Router		/auth/validate_reset_token/ [post]
func (s *ServiceServer) ValidateResetToken(c *fiber.Ctx) error {
	body := new(httprequest.AuthenticationValidateResetToken)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	s.authService.ValidateResetToken(dtoservice.AuthenticationValidateResetTokenReq{
		ResetToken: body.ResetToken,
	})
	return util.NewResponse(c).Success(nil, nil, "Valid")
}
