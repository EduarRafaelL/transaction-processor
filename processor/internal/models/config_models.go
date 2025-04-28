package models

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type EmailConfig struct {
	UrlEmailService string `json:"url_email_service"`
	BodyTemplate    string `json:"body_template"`
	MessageTemplate string `json:"message_template"`
}
