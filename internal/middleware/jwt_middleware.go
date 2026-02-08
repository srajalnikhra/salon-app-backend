package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

func AdminAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"success": false,
				"message": "Missing Authorization header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(
			tokenString,
			&utils.AdminClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(config.LoadJWTConfig().Secret), nil
			},
		)

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"success": false,
				"message": "Invalid or expired token",
			})
		}

		claims := token.Claims.(*utils.AdminClaims)

		// Inject context
		c.Locals("admin_id", claims.AdminID)
		c.Locals("business_id", claims.BusinessID)

		return c.Next()
	}
}
