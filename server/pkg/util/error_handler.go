package util

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Println(err)

	// bad request
	_, ok := err.(BadRequest)
	if ok {
		return NewResponse(c).Error(
			err.Error(),
			MESSAGE_BAD_REQUEST,
			fiber.StatusBadRequest,
		)
	}

	// duplicate
	_, ok = err.(Duplicate)
	if ok {
		return NewResponse(c).Error(
			err.Error(),
			MESSAGE_CONFLICT,
			fiber.StatusConflict,
		)
	}

	// data not found
	_, ok = err.(NoData)
	if ok {
		return NewResponse(c).Error(
			err.Error(),
			MESSAGE_NOT_FOUND,
			fiber.StatusBadRequest,
		)
	}

	// body validation
	tempError, ok := err.(BodyValidationError)
	if ok {
		return NewResponse(c).Error(
			tempError.ListError,
			MESSAGE_FAILED_VALIDATION,
			fiber.StatusBadRequest,
		)
	}

	return NewResponse(c).Error(
		err.Error(),
		MESSAGE_BAD_SYSTEM,
		fiber.StatusBadGateway,
	)
}
