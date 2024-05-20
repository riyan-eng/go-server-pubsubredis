package app

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Put
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Param       body	body	dto.PutExample	true	"body"
// @Router      /example/{id}/ [put]
// @Security ApiKeyAuth
func (s *ServiceServer) PutExample(c *fiber.Ctx) error {
	// parse & validate
	body := new(dto.PutExample)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)

	service := s.exampleService.Put(c.Context(), entity.PutExampleReq{
		UUID:   c.Params("id"),
		Nama:   body.Nama,
		Detail: body.Detail,
	})
	return util.NewResponse(c).Success(service.Data, nil, infrastructure.Localize("OK_UPDATE"))
}
