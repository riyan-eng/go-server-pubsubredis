package util

import (
	"errors"
	"server/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusBadGateway
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	switch err := err.(type) {
	case BadRequest:
		return NewResponse(c).Error(
			err.Error(),
			infrastructure.Localize("BAD_REQUEST"),
			fiber.StatusBadRequest,
		)
	case Duplicate:
		return NewResponse(c).Error(
			err.Error(),
			infrastructure.Localize("CONFLICT"),
			fiber.StatusConflict,
		)
	case NoData:
		return NewResponse(c).Error(
			err.Error(),
			infrastructure.Localize("NOT_FOUND"),
			fiber.StatusBadRequest,
		)
	// case BodyValidationError:
	// 	return NewResponse(c).Error(
	// 		err.ListError,
	// 		infrastructure.Localize("FAILED_VALIDATION"),
	// 		fiber.StatusBadRequest,
	// 	)
	case CustomBadRequest:
		return NewResponse(c).Error(
			err.CustomError(),
			err.Message(),
			err.StatusCode(),
		)
	default:
		return NewResponse(c).Error(
			err.Error(),
			infrastructure.Localize("BAD_SYSTEM"),
			code,
		)
	}
}
