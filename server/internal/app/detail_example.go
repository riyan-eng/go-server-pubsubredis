package app

import (
	"server/infrastructure"
	"server/internal/entity"
	"server/util"

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
	service := s.exampleService.Detail(c.Context(), entity.DetailExampleReq{
		UUID: c.Params("id"),
	})
	return util.NewResponse(c).Success(service.Data, nil, infrastructure.Localize("OK_READ"))
}
