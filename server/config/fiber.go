package config

import (
	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: util.ErrorHandler,
	}
}
