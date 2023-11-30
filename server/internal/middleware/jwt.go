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
			return util.NewResponse(c).Error(nil, "Authorization header is required.", fiber.StatusBadRequest)
		}
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 {
			return util.NewResponse(c).Error(nil, "Undefined token.", fiber.StatusBadRequest)
		}
		tokenString := splitToken[1]
		claim, err := util.ParseToken(tokenString, env.JWT_SECRET_ACCESS)
		if err != nil {
			return util.NewResponse(c).Error(nil, util.MESSAGE_UNAUTHORIZED, fiber.StatusUnauthorized)
		}
		if err := util.ValidateToken(claim, "access"); err != nil {
			return util.NewResponse(c).Error(nil, util.MESSAGE_UNAUTHORIZED, fiber.StatusUnauthorized)
		}
		c.Locals("user_uuid", claim.UserUUID)
		c.Locals("role_code", claim.RoleCode)

		g := new(errgroup.Group)
		g.Go(func() (err error) {
			err = infrastructure.Redis.Expire(c.Context(), fmt.Sprintf("token-%s", claim.UserUUID), time.Minute*env.JWT_EXPIRED_LOGOFF).Err()
			return
		})
		g.Wait()
		return c.Next()
	}
}
