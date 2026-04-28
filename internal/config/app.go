package config

// AppConfig stores application-level configuration
// These settings control the app name, environment, and listening port
type AppConfig struct {
	Name string // Application name
	Env  string // Environment (development, production, etc.)
	Port string // Server port
}

// LoadAppConfig loads app configuration from environment variables
func LoadAppConfig() AppConfig {
	return AppConfig{
		Name: GetEnv("APP_NAME", "Salon Backend"),
		Env:  GetEnv("APP_ENV", "development"),
		Port: GetEnv("APP_PORT", "3000"),
	}
}
