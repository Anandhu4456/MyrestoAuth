package cfg

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	JWTAccessExpiryMinute int
	JWTAccessSecret       string
	JWTRefreshExpiryDays  int
	JWTRefreshSecret      string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("env file loading failed due to : %w", err)
	}

	cfg := &Config{
		DBHost:     env("DB_HOST", "localhost"),
		DBPort:     env("DB_PORT", "5432"),
		DBUser:     env("DB_USER", "postgres"),
		DBPassword: env("DB_PASSWORD", ""),
		DBName:     env("DB_NAME", "myrestodb"),
		DBSSLMode:  env("DB_SSLMODE", "disable"),

		JWTAccessExpiryMinute: getEnvAsInt("JWT_ACCESS_EXPIRY_MINUTE", 15),
		JWTAccessSecret:       env("JWT_ACCESS_SECRET", ""),
		JWTRefreshExpiryDays:  getEnvAsInt("JWT_REFRESH_EXPIRY_DAYS", 7),
		JWTRefreshSecret:      env("JWT_REFRESH_SECRET", ""),
	}
	return cfg, nil
}

func env(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok {
		envInt, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("Invalid value for %s, using default: %d", key, fallback)
			return fallback
		}
		return envInt
	}
	return fallback
}
