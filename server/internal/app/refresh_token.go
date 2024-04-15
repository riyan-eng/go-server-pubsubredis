package app

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Refresh Token
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthenticationRefreshToken	true  "body"
// @Router		/auth/refresh_token/ [post]
func (s *ServiceServer) RefreshToken(c *fiber.Ctx) error {
	body := new(dto.AuthenticationRefreshToken)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("OK_CREATE"), fiber.StatusBadRequest)
	}

	service := s.authService.RefreshToken(entity.AuthenticationRefreshTokenReq{
		RefreshToken: body.RefreshToken,
		Issuer:       string(c.Request().Host()),
	})
	data := fiber.Map{
		"access_token":  service.AccessToken,
		"refresh_token": service.RefreshToken,
		"expired_at":    service.ExpiredAt.Time,
	}
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"))
}
