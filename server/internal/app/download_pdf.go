package app

import (
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Example
// @Tags       	PDF
// @Router      /pdf/	[get]
// @Security ApiKeyAuth
func (s *ServiceServer) DownloadPDF(c *fiber.Ctx) error {
	pdf := s.exampleService.Pdf(c.Context())
	fileName := "pdf_example"
	c.Response().Header.Set("Content-Type", "application/pdf")
	c.Response().Header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v.pdf", fileName))
	io.Copy(c.Response().BodyWriter(), pdf.Buffer())
	return nil
}
