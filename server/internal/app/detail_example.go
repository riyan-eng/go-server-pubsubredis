package app

import (
	dtoservice "server/internal/dto_service"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Detail
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Router      /example/{id}/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) DetailExample(c *fiber.Ctx) error {
	service := s.exampleService.Detail(dtoservice.DetailExampleReq{
		UUID: util.NewQuery().CheckExistingData("example", "example", c.Params("id")),
	})
	return util.NewResponse(c).Success(service.Item, nil, util.MESSAGE_OK_READ)
}
