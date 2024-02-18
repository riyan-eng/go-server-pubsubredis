package config

import (
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func NewFiberConfig() fiber.Config {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return fiber.Config{
		ErrorHandler:  util.ErrorHandler,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "Test App v1.0.1",
		Concurrency:   1024 * 1024,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	}
}
