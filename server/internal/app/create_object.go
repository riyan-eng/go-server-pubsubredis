package app

import (
	dtoservice "server/internal/dto_service"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Create
// @Tags        Object
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.CreateExample	true  "body"
// @Router		/object/ [post]
// @Security ApiKeyAuth
func (s *ServiceServer) CreateObject(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	util.PanicIfNeeded(err)
	fileMeta := util.NewFile().SaveLocal(c, file, "default_private")
	s.objectService.Create(dtoservice.CreateObjectReq{
		UUID:     fileMeta.UUID,
		Bukcet:   "default_private",
		Nama:     fileMeta.Nama,
		Size:     fileMeta.Size,
		MimeType: fileMeta.MimeType,
		Url:      fileMeta.Url,
		Path:     fileMeta.Path,
	})
	data := fiber.Map{
		"url": fileMeta.Url,
	}
	return util.NewResponse(c).Success(data, nil, util.MESSAGE_OK_CREATE)
}
