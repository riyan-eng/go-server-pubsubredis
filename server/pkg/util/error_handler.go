package util

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Println(err)

	// bad request
	_, ok := err.(BadRequest)
	if ok {
		return NewResponse(c).Error(err.Error(), MESSAGE_BAD_REQUEST, fiber.StatusBadRequest)
	}

	// validation
	_, ok = err.(ValidationError)
	if ok {

		return NewResponse(c).Error(err.Error(), MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	// no data
	if err == sql.ErrNoRows {
		return NewResponse(c).Error(err.Error(), MESSAGE_NOT_FOUND, fiber.StatusBadRequest)
	}

	// duplicate
	// if StringNumToInt(fmt.Sprintf("%v", err.(*pq.Error).Code)) == 23505 {
	// 	return NewResponse(c).Error(err.Error(), MESSAGE_CONFLICT, fiber.StatusConflict)
	// }

	return NewResponse(c).Error(err.Error(), MESSAGE_BAD_SYSTEM, fiber.StatusBadGateway)
}
