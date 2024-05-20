package util

import (
	"errors"
	"server/infrastructure"

	"github.com/gofiber/fiber/v2"
)

type CustomBadRequest struct {
	temp        string
	Errors      any
	Messages    string
	StatusCodes int
}

func (validationError CustomBadRequest) Error() string {
	return validationError.temp
}

func (validationError CustomBadRequest) CustomError() any {
	return validationError.Errors
}

var localizeResponseCode = map[int]string{
	400: "BAD_REQUEST",
	409: "CONFLICT",
	404: "NOT_FOUND",
	500: "BAD_SYSTEM",
}

func (validationError CustomBadRequest) Message() string {
	if validationError.Messages == "" {
		validationError.Messages = infrastructure.Localize(localizeResponseCode[validationError.StatusCodes])
	}
	return validationError.Messages
}

func (validationError CustomBadRequest) StatusCode() int {
	if validationError.StatusCodes == 0 {
		validationError.StatusCodes = 400
	}
	return validationError.StatusCodes
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusBadGateway
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	switch err := err.(type) {
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
