package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

func Auth() fiber.Handler {
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
			&utils.Claims{},
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

		claims := token.Claims.(*utils.Claims)

		c.Locals("business_id", claims.BusinessID)
		c.Locals("role", claims.Role)

		if claims.AdminID != nil {
			c.Locals("admin_id", *claims.AdminID)
		}
		if claims.StaffID != nil {
			c.Locals("staff_id", *claims.StaffID)
		}

		return c.Next()
	}
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Locals("role") != "admin" {
			return c.Status(403).JSON(fiber.Map{
				"success": false,
				"message": "Admin access required",
			})
		}
		return c.Next()
	}
}
