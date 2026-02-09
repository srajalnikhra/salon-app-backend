package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
)

type Claims struct {
	AdminID    *uint  `json:"admin_id,omitempty"`
	StaffID    *uint  `json:"staff_id,omitempty"`
	BusinessID uint   `json:"business_id"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

func generateToken(claims Claims) (string, error) {
	jwtConfig := config.LoadJWTConfig()

	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(
			time.Now().Add(time.Duration(jwtConfig.ExpiryDays) * 24 * time.Hour),
		),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

func GenerateAdminToken(adminID, businessID uint) (string, error) {
	return generateToken(Claims{
		AdminID:    &adminID,
		BusinessID: businessID,
		Role:       "admin",
	})
}

func GenerateStaffToken(staffID, businessID uint) (string, error) {
	return generateToken(Claims{
		StaffID:    &staffID,
		BusinessID: businessID,
		Role:       "staff",
	})
}
