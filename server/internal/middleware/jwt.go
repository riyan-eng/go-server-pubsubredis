package middleware

import (
	"strings"

	"server/util"

	"github.com/gofiber/fiber/v2"
)

func AuthorizeJwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			util.PanicIfNeeded(util.CustomBadRequest{
				Messages:    "Authorization header is required.",
				StatusCodes: 400,
			})
		}

		token, found := strings.CutPrefix(authHeader, "Bearer ")
		if !found {
			util.PanicIfNeeded(util.CustomBadRequest{
				Messages:    "Undefined token.",
				StatusCodes: 400,
			})
		}

		claimss, err := util.NewToken().ParseAccess(token)
		if err != nil {
			util.PanicIfNeeded(util.CustomBadRequest{
				Errors:      err.Error(),
				Messages:    "Unauthorized.",
				StatusCodes: 401,
			})
		}

		if err := util.NewToken().ValidateAccess(c.Context(), claimss); err != nil {
			util.PanicIfNeeded(util.CustomBadRequest{
				Errors:      err.Error(),
				Messages:    "Unauthorized.",
				StatusCodes: 401,
			})
		}
		c.Locals("claims", claimss)

		// g := new(errgroup.Group)
		// g.Go(func() (err error) {
		// 	err = infrastructure.Redis.Expire(c.Context(), fmt.Sprintf("token-%s", claim.UserUUID), time.Minute*env.NewEnvironment().JWT_EXPIRED_LOGOFF).Err()
		// 	return
		// })
		// g.Wait()
		return c.Next()
	}
}
