package app

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/internal/entity"
	"server/util"

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
	queryParam := new(dto.PaginationReq).Init()
	err := c.QueryParser(&queryParam)
	util.PanicIfNeeded(err)

	pageMeta := util.NewPagination().GetPageMeta(queryParam.Page, queryParam.Limit)

	service := s.exampleService.List(entity.ListExampleReq{
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

	return util.NewResponse(c).Success(service.DataData, meta, infrastructure.Localize("OK_CREATE"))
}
