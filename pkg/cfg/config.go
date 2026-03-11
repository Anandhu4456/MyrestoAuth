package cfg

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL string

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

	SMTPHost        string
	SMTPPort        int
	SMTPEmail       string
	SMTPAppPassword string
	EmailFromName   string

	VerificationTokenExpiryHours int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("env file loading failed due to : %v", err)
	}

	cfg := &Config{

		BaseURL: env("BASE_URL", "https://myrestotody.com"),

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

		SMTPHost:        env("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:        getEnvAsInt("SMTP_PORT", 587),
		SMTPEmail:       env("SMTP_EMAIL", ""),
		SMTPAppPassword: env("SMTP_APP_PASSWORD", ""),
		EmailFromName:   env("EMAIL_FROM_NAME", "MyRestoToday"),

		VerificationTokenExpiryHours: getEnvAsInt("VERIFICATION_TOKEN_EXPIRY_HOURS", 24),
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
