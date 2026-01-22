package config

import "strconv"

type JWTConfig struct {
	Secret     string
	ExpiryDays int
}

func LoadJWTConfig() JWTConfig {
	days, _ := strconv.Atoi(GetEnv("JWT_EXPIRY_DAYS", "7"))

	return JWTConfig{
		Secret:     GetEnv("JWT_SECRET", "secret"),
		ExpiryDays: days,
	}
}
