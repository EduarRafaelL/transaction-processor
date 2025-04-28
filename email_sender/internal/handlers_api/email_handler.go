package handlersInternal

import (
	"email-sender/internal/models"
	"email-sender/internal/services"
	"encoding/json"
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
	// Validar que sea método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parsear el body del request
	var email models.GenericEmail
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Enviar el correo
	err := h.EmailService.SendEmail(email)
	if err != nil {
		http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Email sent successfully",
	}
	json.NewEncoder(w).Encode(response)
}
