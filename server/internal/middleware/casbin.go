package middleware

import (
	"server/pkg/util"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func PermitCasbin(enforce *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get current user
		userID, ok := c.Locals("user_id").(string)
		if userID == "" || !ok {
			return util.NewResponse(c).Error(
				"Current logged in user not found.",
				util.MESSAGE_UNAUTHORIZED, fiber.StatusUnauthorized,
			)
		}

		// load new change policy
		if err := enforce.LoadPolicy(); err != nil {
			return util.NewResponse(c).Error(
				"Failed to load casbin policy.",
				util.MESSAGE_BAD_SYSTEM,
				fiber.StatusInternalServerError,
			)
		}
		// casbin enforce policy
		accepted, err := enforce.Enforce(userID, c.OriginalURL(), c.Method()) // userID - url - method
		if err != nil {
			return util.NewResponse(c).Error(
				err.Error(),
				"Error when authorizing user's accessibility.",
				fiber.StatusBadRequest,
			)
		}
		if !accepted {
			return util.NewResponse(c).Error(
				nil,
				"You are not allowed.",
				fiber.StatusForbidden,
			)
		}
		return c.Next()
	}
}
