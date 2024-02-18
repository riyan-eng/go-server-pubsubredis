package app

// @Summary     Detail
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Router      /object/{id}/ [get]
// @Security ApiKeyAuth
// func (s *ServiceServer) DetailObject(c *fiber.Ctx) error {
// 	service := s.objectService.Detail(entity.DetailObjectReq{
// 		ID: util.NewQuery().GetIDByUUID("objects", c.Params("id")),
// 	})
// 	return util.NewResponse(c).Success(service.Item, nil, util.MESSAGE_OK_READ)
// }
