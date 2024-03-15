package middleware

import (
	"fmt"
	"strings"
	"time"

	"server/env"
	"server/infrastructure"

	"server/pkg/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
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
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 {
			util.PanicIfNeeded(util.CustomBadRequest{
				Messages:    "Undefined token.",
				StatusCodes: 400,
			})
		}
		tokenString := splitToken[1]
		claim, err := util.ParseToken(tokenString, env.NewEnvironment().JWT_SECRET_ACCESS)
		if err != nil {
			util.PanicIfNeeded(util.CustomBadRequest{
				Errors:      err.Error(),
				Messages:    "Unauthorized.",
				StatusCodes: 401,
			})
		}
		if err := util.ValidateToken(claim, "access"); err != nil {
			util.PanicIfNeeded(util.CustomBadRequest{
				Errors:      err.Error(),
				Messages:    "Unauthorized.",
				StatusCodes: 401,
			})
		}
		c.Locals("user_uuid", claim.UserUUID)
		c.Locals("role_code", claim.RoleCode)

		g := new(errgroup.Group)
		g.Go(func() (err error) {
			err = infrastructure.Redis.Expire(c.Context(), fmt.Sprintf("token-%s", claim.UserUUID), time.Minute*env.NewEnvironment().JWT_EXPIRED_LOGOFF).Err()
			return
		})
		g.Wait()
		return c.Next()
	}
}
