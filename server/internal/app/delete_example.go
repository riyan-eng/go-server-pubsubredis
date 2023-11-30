package app

import (
	dtoservice "server/internal/dto_service"
	"server/pkg/util"

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
	s.exampleService.Delete(dtoservice.DeleteExampleReq{
		UUID: util.NewQuery().CheckExistingData("example", "example", c.Params("id")),
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_DELETE)
}
