package clients

import (
	"email-sender/internal/models"

	"gopkg.in/gomail.v2"
)

func NewEmailSender(config models.EmailConfig) *gomail.Dialer {
	dialer := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUsername, config.SMTPPassword)
	return dialer
}
