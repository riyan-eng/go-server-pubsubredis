package app

import (
	"server/internal/dto"
	"server/internal/entity"
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

func (s *ServiceServer) ImportExample(c *fiber.Ctx) error {
	file, err := c.FormFile("import_file")
	if err != nil {
		util.PanicIfNeeded(util.BadRequest{Message: "tidak ada file yang diunggah terkait dengan key yang diberikan"})
	}
	body := util.ReadImportExcel[[]dto.ImportExample](file)
	var items []entity.ImportExampleItemReq
	for _, i := range body {
		items = append(items, entity.ImportExampleItemReq{
			Nama:   i.Nama,
			Detail: i.Detail,
		})
	}

	s.exampleService.Import(entity.ImportExampleReq{
		Items: items,
	})

	return nil
}
