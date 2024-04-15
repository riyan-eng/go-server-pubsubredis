package app

import (
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Delete
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Router      /example/{id}/ [delete]
// @Security ApiKeyAuth
func (s *ServiceServer) DeleteExample(c *fiber.Ctx) error {
	s.exampleService.Delete(entity.DeleteExampleReq{
		UUID: util.NewQuery().CheckExistingData("example", "example", c.Params("id")),
	})
	return util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_CREATE"))
}
