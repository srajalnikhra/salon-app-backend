package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

// Auth is middleware that validates JWT tokens and extracts user information
// Intercepts incoming requests and checks for valid Bearer token in Authorization header
// Extracts user claims and stores them in request context for downstream handlers
func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve Authorization header from incoming request
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"success": false,
				"message": "Missing Authorization header",
			})
		}

		// Extract token from Bearer {token} format
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate JWT token using secret key
		token, err := jwt.ParseWithClaims(
			tokenString,
			&utils.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(config.LoadJWTConfig().Secret), nil
			},
		)

		// Return 401 if token is invalid or expired
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"success": false,
				"message": "Invalid or expired token",
			})
		}

		// Extract claims from validated token
		claims := token.Claims.(*utils.Claims)

		// Store user information in request context for downstream handlers
		c.Locals("business_id", claims.BusinessID)
		c.Locals("role", claims.Role)

		// Store admin or staff ID if present
		if claims.AdminID != nil {
			c.Locals("admin_id", *claims.AdminID)
		}
		if claims.StaffID != nil {
			c.Locals("staff_id", *claims.StaffID)
		}

		return c.Next()
	}
}

// AdminOnly is middleware that checks if user has admin role
// Must be applied after Auth middleware to ensure token validation
// Blocks access for non-admin users (staff members)
func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if user role stored by Auth middleware is admin
		if c.Locals("role") != "admin" {
			return c.Status(403).JSON(fiber.Map{
				"success": false,
				"message": "Admin access required",
			})
		}
		// Proceed to next handler if admin
		return c.Next()
	}
}
