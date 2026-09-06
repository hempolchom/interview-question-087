package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPass             string
	DBName             string
	JWTSecret          string
	JWTExpirationHours int
}

var AppConfig *Config

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found")
	}

	jwtExpHours, _ := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24"))

	AppConfig = &Config{
		Port:               getEnv("PORT", "9087"),
		DBHost:             getEnv("DB_HOST", "postgres-db"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPass:             getEnv("DB_PASSWORD", "12345"),
		DBName:             getEnv("DB_NAME", "interview_db"),
		JWTSecret:          getEnv("JWT_SECRET", "supersecretkey_087"),
		JWTExpirationHours: jwtExpHours,
	}

	return AppConfig
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName)
}