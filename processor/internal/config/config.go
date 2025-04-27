package config

import (
	"os"
	"transaction-processor/internal/models"
)

type Config struct {
	InputFilePath string             `json:"input_file_path"`
	Delimiter     string             `json:"delimiter"`
	DBConfig      models.DBConfig    `json:"db_config"`
	EmailConfig   models.EmailConfig `json:"email_config"`
}

func LoadConfig() (*Config, error) {

	config := &Config{
		InputFilePath: os.Getenv("INPUT_FILE_PATH"),
		Delimiter:     os.Getenv("DELIMITER"),
		DBConfig: models.DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
		EmailConfig: models.EmailConfig{
			UrlEmailService: os.Getenv("EMAIL_SERVICE_URL"),
			BodyTemplate:    os.Getenv("EMAIL_BODY_TEMPLATE"),
			MessageTemplate: os.Getenv("EMAIL_MESSAGE_TEMPLATE"),
			From:            os.Getenv("EMAIL_FROM"),
		},
	}

	return config, nil
}
