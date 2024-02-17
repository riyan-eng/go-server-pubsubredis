package app

import (
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     View
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       bucket	path	string				true	"bucket"
// @Param       id		path	string				true	"id"
// @Router      /object/{bucket}/{id}/view [get]
// @Security ApiKeyAuth
func (s *ServiceServer) ViewObject(c *fiber.Ctx) error {
	service := s.objectService.Detail(entity.DetailObjectReq{
		ID: util.NewQuery().GetIDByUUID("objects", c.Params("id")),
	})
	c.Response().Header.SetContentType(service.Item.MimeType)
	return c.SendFile(service.Item.Path)
}
