package app

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Create
// @Tags        Example
// @Accept		json
// @Produce		json
// @Param       body	body  dto.CreateExample	true  "body"
// @Router		/example/ [post]
// @Security ApiKeyAuth
func (s *ServiceServer) CreateExample(c *fiber.Ctx) error {
	body := new(dto.CreateExample)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	errors, err := util.NewValidation().ValidateStruct(*body)
	util.PanicBodyValidation(errors, err)

	service := s.exampleService.Create(c.Context(), entity.CreateExampleReq{
		Nama:   body.Nama,
		Detail: body.Detail,
	})
	return util.NewResponse(c).Success(service.Item, nil, infrastructure.Localize("OK_CREATE"), 201)
}
