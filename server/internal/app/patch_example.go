package app

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Put
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Param       body	body	dto.PatchExample	true	"body"
// @Router      /example/{id}/ [patch]
// @Security ApiKeyAuth
func (s *ServiceServer) PatchExample(c *fiber.Ctx) error {
	// parse & validate
	body := new(dto.PatchExample)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)

	s.exampleService.Patch(entity.PatchExampleReq{
		UUID:   util.NewQuery().CheckExistingData("example", "example", c.Params("id")),
		Nama:   body.Nama,
		Detail: body.Detail,
	})
	return util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_CREATE"))
}
