package services

import (
	"bytes"
	"email-sender/internal/models"
	"email-sender/internal/repositories"
	"fmt"
	"html/template"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	dialer *gomail.Dialer
	repo   *repositories.TemplateRepository
}

func NewEmailService(dialer *gomail.Dialer, repo *repositories.TemplateRepository) *EmailService {
	return &EmailService{
		dialer: dialer,
		repo:   repo,
	}
}

func (s *EmailService) SendEmail(email models.GenericEmail) error {
	content, err := s.GetEmailContent(email)
	if err != nil {
		return fmt.Errorf("error getting email content: %w", err)
	}

	mailer := s.createMailer(email, content)
	if err := s.dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}
	return nil
}

func (s *EmailService) createMailer(email models.GenericEmail, emailContent string) *gomail.Message {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.dialer.Username)
	mailer.SetHeader("To", email.To)
	mailer.SetHeader("Subject", email.Subject)
	mailer.SetBody("text/html", emailContent)
	return mailer
}

func (s *EmailService) GetEmailContent(email models.GenericEmail) (string, error) {
	bodyTemplateContent, err := s.repo.GetTemplateByName(email.BodyTemplate)
	if err != nil {
		return "", fmt.Errorf("error loading body template: %w", err)
	}
	messageTemplateContent, err := s.repo.GetTemplateByName(email.MessageTemplate)
	if err != nil {
		return "", fmt.Errorf("error loading message template: %w", err)
	}

	var messageBuffer bytes.Buffer
	messageTmpl, err := template.New("messageTemplate").Parse(messageTemplateContent.HtmlBody)
	if err != nil {
		return "", fmt.Errorf("error parsing message template: %w", err)
	}
	err = messageTmpl.Execute(&messageBuffer, email.Body)
	if err != nil {
		return "", fmt.Errorf("error executing message template: %w", err)
	}

	renderedBody := messageBuffer.String()
	templateData := map[string]any{
		"body": template.HTML(renderedBody),
	}

	var bodyBuffer bytes.Buffer
	bodyTmpl, err := template.New("bodyTemplate").Parse(bodyTemplateContent.HtmlBody)
	if err != nil {
		return "", fmt.Errorf("error parsing body template: %w", err)
	}
	err = bodyTmpl.Execute(&bodyBuffer, templateData)
	if err != nil {
		return "", fmt.Errorf("error executing body template: %w", err)
	}
	return bodyBuffer.String(), nil
}
