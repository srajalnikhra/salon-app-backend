package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
)

// Claims extends JWT standard claims with custom user information
// Stores user type (admin/staff), IDs, and business context
type Claims struct {
	AdminID              *uint  `json:"admin_id,omitempty"` // Admin ID if user is admin
	StaffID              *uint  `json:"staff_id,omitempty"` // Staff ID if user is staff
	BusinessID           uint   `json:"business_id"`        // Salon/business context
	Role                 string `json:"role"`               // User role (admin or staff)
	jwt.RegisteredClaims        // Standard JWT claims
}

// generateToken creates a signed JWT token with given claims
// Token includes expiration time from config
func generateToken(claims Claims) (string, error) {
	jwtConfig := config.LoadJWTConfig()

	// Set token expiration time based on configuration
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(
			time.Now().Add(time.Duration(jwtConfig.ExpiryDays) * 24 * time.Hour),
		),
	}

	// Create JWT token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign token with secret key and return as string
	return token.SignedString([]byte(jwtConfig.Secret))
}

// GenerateAdminToken creates JWT token for admin user
// Token includes admin ID, business ID, and admin role
func GenerateAdminToken(adminID, businessID uint) (string, error) {
	return generateToken(Claims{
		AdminID:    &adminID,
		BusinessID: businessID,
		Role:       "admin",
	})
}

// GenerateStaffToken creates JWT token for staff user
// Token includes staff ID, business ID, and staff role
func GenerateStaffToken(staffID, businessID uint) (string, error) {
	return generateToken(Claims{
		StaffID:    &staffID,
		BusinessID: businessID,
		Role:       "staff",
	})
}
