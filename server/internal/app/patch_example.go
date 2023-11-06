package app

import (
	dtoservice "server/internal/dto_service"
	httprequest "server/pkg/http.request"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Put
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Param       body	body	httprequest.PatchExample	true	"body"
// @Router      /example/{id}/ [patch]
// @Security ApiKeyAuth
func (s *ServiceServer) PatchExample(c *fiber.Ctx) error {
	// parse & validate
	body := new(httprequest.PatchExample)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)

	s.exampleService.Patch(dtoservice.PatchExampleReq{
		ID:     util.NewQuery().GetIDByUUID("example", c.Params("id")),
		Nama:   body.Nama,
		Detail: body.Detail,
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_UPDATE)
}
