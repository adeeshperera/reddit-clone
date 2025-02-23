package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
	}
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
}

func Load() *Config {
	// Load .env file
	godotenv.Load()

	cfg := &Config{}

	// Server configuration
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	// PostgreSQL configuration
	cfg.Postgres.Host = getEnv("POSTGRES_HOST", "localhost")
	cfg.Postgres.Port = getEnv("POSTGRES_PORT", "5432")
	cfg.Postgres.User = getEnv("POSTGRES_USER", "postgres")
	cfg.Postgres.Password = getEnv("POSTGRES_PASSWORD", "postgres")
	cfg.Postgres.DBName = getEnv("POSTGRES_DB", "myapp")

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
