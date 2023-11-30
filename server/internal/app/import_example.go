package app

import (
	dtoservice "server/internal/dto_service"
	httprequest "server/pkg/http.request"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

func (s *ServiceServer) ImportExample(c *fiber.Ctx) error {
	file, err := c.FormFile("import_file")
	if err != nil {
		util.PanicIfNeeded(util.BadRequest{Message: "tidak ada file yang diunggah terkait dengan key yang diberikan"})
	}
	body := util.ReadImportExcel[[]httprequest.ImportExample](file)
	var items []dtoservice.ImportExampleItemReq
	for _, i := range body {
		items = append(items, dtoservice.ImportExampleItemReq{
			Nama:   i.Nama,
			Detail: i.Detail,
		})
	}

	s.exampleService.Import(dtoservice.ImportExampleReq{
		Items: items,
	})

	return nil
}
