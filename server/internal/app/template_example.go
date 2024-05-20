package app

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary     Example
// @Tags       	Template
// @Accept		json
// @Produce		json
// @Router      /template/example/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) TemplateExample(c *fiber.Ctx) error {
	service := s.exampleService.Template(c.Context())
	c.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=template_example.xlsx")
	return service.File.Write(c)
}
