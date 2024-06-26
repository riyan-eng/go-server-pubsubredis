package app

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Register
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthenticationRegister	true  "body"
// @Router		/auth/register/ [post]
func (s *ServiceServer) Register(c *fiber.Ctx) error {
	body := new(dto.AuthenticationRegister)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("OK_CREATE"), fiber.StatusBadRequest)
	}

	s.authService.Register(entity.AuthenticationRegisterReq{
		UUIDUser:     uuid.NewString(),
		UUIDUserData: uuid.NewString(),
		Nama:         body.Nama,
		Email:        body.Email,
		Password:     body.Password,
		KodeRole:     "MASYARAKAT",
		NIK:          body.NIK,
		NomorTelepon: body.NomorTelepon,
	})
	return util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_CREATE"))
}
