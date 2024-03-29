package app

// @Summary     Download
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       bucket	path	string				true	"bucket"
// @Param       id		path	string				true	"id"
// @Router      /object/{bucket}/{id}/download [get]
// @Security ApiKeyAuth
// func (s *ServiceServer) DownloadObject(c *fiber.Ctx) error {
// 	service := s.objectService.Detail(entity.DetailObjectReq{
// 		ID: util.NewQuery().GetIDByUUID("objects", c.Params("id")),
// 	})
// 	c.Response().Header.SetContentType("application/octet-stream")
// 	c.Response().Header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v", service.Item.Nama))
// 	return c.SendFile(service.Item.Path)
// }
