package config

import "strconv"

// JWTConfig stores JWT token generation and validation settings
type JWTConfig struct {
	Secret     string // Secret key for signing JWT tokens
	ExpiryDays int    // Token expiration time in days
}

// LoadJWTConfig loads JWT configuration from environment variables
func LoadJWTConfig() JWTConfig {
	days, _ := strconv.Atoi(GetEnv("JWT_EXPIRY_DAYS", "7"))

	return JWTConfig{
		Secret:     GetEnv("JWT_SECRET", "secret"),
		ExpiryDays: days,
	}
}
