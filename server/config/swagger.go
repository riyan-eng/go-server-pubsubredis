package config

import (
	"html/template"

	"github.com/gofiber/swagger"
)

func NewSwaggerConfig() swagger.Config {
	return swagger.Config{
		Title:  "Golang",
		Layout: "StandaloneLayout",
		Plugins: []template.JS{
			template.JS("SwaggerUIBundle.plugins.DownloadUrl"),
		},
		Presets: []template.JS{
			template.JS("SwaggerUIBundle.presets.apis"),
			template.JS("SwaggerUIStandalonePreset"),
		},
		DeepLinking:              true,
		DefaultModelsExpandDepth: 1,
		DefaultModelExpandDepth:  1,
		DefaultModelRendering:    "example",
		DocExpansion:             "list",
		ShowMutatedRequest:       true,
	}
}
