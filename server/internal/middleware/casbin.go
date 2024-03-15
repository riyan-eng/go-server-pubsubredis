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
			util.PanicIfNeeded(util.CustomBadRequest{
				Messages:    "Current logged in user not found.",
				StatusCodes: 401,
			})
		}

		// load new change policy
		if err := enforce.LoadPolicy(); err != nil {
			util.PanicIfNeeded(util.CustomBadRequest{
				Errors:      err.Error(),
				Messages:    "Failed to load policy.",
				StatusCodes: 500,
			})
		}
		// casbin enforce policy
		accepted, err := enforce.Enforce(userID, c.OriginalURL(), c.Method()) // userID - url - method
		if err != nil {
			util.PanicIfNeeded(util.CustomBadRequest{
				Errors:      err.Error(),
				Messages:    "Error when authorizing user's accessibility.",
				StatusCodes: 400,
			})
		}
		if !accepted {
			util.PanicIfNeeded(util.CustomBadRequest{
				Messages:    "You are not allowed.",
				StatusCodes: 403,
			})
		}
		return c.Next()
	}
}
