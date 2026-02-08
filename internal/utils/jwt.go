package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
)

type AdminClaims struct {
	AdminID    uint `json:"admin_id"`
	BusinessID uint `json:"business_id"`
	jwt.RegisteredClaims
}

func GenerateAdminToken(adminID, businessID uint) (string, error) {
	jwtConfig := config.LoadJWTConfig()

	claims := AdminClaims{
		AdminID:    adminID,
		BusinessID: businessID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Duration(jwtConfig.ExpiryDays) * 24 * time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}
