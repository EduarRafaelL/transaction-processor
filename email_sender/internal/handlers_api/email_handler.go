package handlersInternal

import (
	"email-sender/internal/services"
	"net/http"
)

type EmailHandler struct {
	EmailService services.EmailService
}

func NewEmailHandler(emailService services.EmailService) *EmailHandler {
	return &EmailHandler{
		EmailService: emailService,
	}
}

func (h *EmailHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	//Implement the logic to handle the email sending requests
}
