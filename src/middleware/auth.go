package middleware

import (
	"go-fiber/src/config"
	"go-fiber/src/lib"
	"go-fiber/src/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware validates JWT tokens
func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return utils.ErrorResponse(c, 401, config.NotFoundCode, "Authorization header is required")
	}

	claims, err := lib.ValidateToken(token)
	if err != nil {
		return utils.ErrorResponse(c, 401, config.UnauthorizedCode, "Invalid or expired token")
	}

	if claims["type"] != "access_token" {
		return utils.ErrorResponse(c, 403, config.ForbiddenCode, "Access denied")
	}

	return c.Next()
}
