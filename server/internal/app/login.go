package app

import (
	"server/internal/dto"
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Login
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthenticationLogin	true  "body"
// @Router		/auth/login/ [post]
func (s *ServiceServer) Login(c *fiber.Ctx) error {
	body := new(dto.AuthenticationLogin)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	service := s.authService.Login(entity.AuthenticationLoginReq{
		Email:    body.Email,
		Password: body.Password,
		Issuer:   string(c.Request().Host()),
	})
	if !service.Match {
		return util.NewResponse(c).Error(nil, util.MESSAGE_FAILED_LOGIN, fiber.StatusBadRequest)
	}
	data := fiber.Map{
		"access_token":  service.AccessToken,
		"refresh_token": service.RefreshToken,
		"expired_at":    service.ExpiredAt.Time,
	}
	return util.NewResponse(c).Success(data, nil, util.MESSAGE_OK_LOGIN)
}
