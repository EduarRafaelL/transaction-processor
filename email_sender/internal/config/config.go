package config

import (
	"email-sender/internal/models"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBConfig    models.DBConfig    `json:"db_config"`
	EmailConfig models.EmailConfig `json:"email_config"`
}

func LoadConfig() (*Config, error) {

	config := &Config{
		DBConfig: models.DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
		EmailConfig: models.EmailConfig{
			SMTPHost:     os.Getenv("SMTP_HOST"),
			SMTPPort:     parseEnvToInt("SMTP_PORT"),
			SMTPUsername: os.Getenv("SMTP_USERNAME"),
			SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		},
	}

	return config, nil
}

func parseEnvToInt(key string) int {
	valueStr := os.Getenv(key)
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Error parsing env variable %s: %v\n", key, err)
		return 0
	}
	return valueInt
}
