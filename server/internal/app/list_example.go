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

	service := s.exampleService.List(c.Context(), entity.ListExampleReq{
		Search: queryParam.Search,
		Order:  queryParam.Order,
		Limit:  pageMeta.Limit,
		Offset: pageMeta.Offset,
	})

	meta := util.PaginationMeta{
		Page:       pageMeta.Page,
		Limit:      pageMeta.Limit,
		CountRows:  service.CountRows,
		CountPages: util.NewPagination().GetCountPages(service.CountRows, pageMeta.Limit),
	}

	return util.NewResponse(c).Success(service.Data, meta, infrastructure.Localize("OK_READ"))

	// accessToken, accessExpiredAt, err := util.NewToken().CreateAccess(c.Context(), "id")
	// if err != nil {
	// 	return util.NewResponse(c).Error(err, "token error", 400)
	// }

	// refreshToken, refreshExpiredAt, err := util.NewToken().CreateRefresh(c.Context(), "id")
	// if err != nil {
	// 	return util.NewResponse(c).Error(err, "token error", 400)
	// }

	// data := fiber.Map{
	// 	"access_token":       accessToken,
	// 	"access_expired_at":  accessExpiredAt.Local(),
	// 	"refresh_token":      refreshToken,
	// 	"refresh_expired_at": refreshExpiredAt.Local(),
	// }
	// return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"))

}
