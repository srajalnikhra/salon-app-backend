package config

// DBConfig stores database connection configuration
// These settings are used to establish a connection to PostgreSQL
type DBConfig struct {
	Host     string // Database host address
	Port     string // Database port
	User     string // Database user
	Password string // Database password
	Name     string // Database name
}

// LoadDBConfig loads database configuration from environment variables
func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     GetEnv("DB_HOST", "localhost"),
		Port:     GetEnv("DB_PORT", "5432"),
		User:     GetEnv("DB_USER", "postgres"),
		Password: GetEnv("DB_PASSWORD", ""),
		Name:     GetEnv("DB_NAME", "salon_app"),
	}
}
