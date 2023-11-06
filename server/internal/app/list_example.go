package app

import (
	dtoservice "server/internal/dto_service"
	httprequest "server/pkg/http.request"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

// @Summary      List
// @Tags       	 Example
// @Produce      json
// @Param        order		query   string	false  "desc/asc default(desc)"
// @Param        search		query   string	false  "search"
// @Param        page		query   int		false  "page"
// @Param        per_page	query   int		false  "per_page"
// @Router       /example/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) ListExample(c *fiber.Ctx) error {
	queryParam := new(httprequest.PaginationReq).Init()
	err := c.QueryParser(&queryParam)
	util.PanicIfNeeded(err)

	pageMeta := util.NewPagination().GetPageMeta(queryParam.Page, queryParam.Limit)

	service := s.exampleService.List(dtoservice.ListExampleReq{
		Search: queryParam.Search,
		Order:  queryParam.Order,
		Limit:  pageMeta.Limit,
		Offset: pageMeta.Offset,
	})

	meta := util.PaginationMeta{
		Page:       pageMeta.Page,
		Limit:      pageMeta.Limit,
		TotalRows:  service.Total,
		TotalPages: util.NewPagination().GetTotalPages(service.Total, pageMeta.Limit),
	}

	return util.NewResponse(c).Success(service.Items, meta, util.MESSAGE_OK_READ)
}
