package main

import (
	"email-sender/internal/clients"
	"email-sender/internal/config"
	handlersInternal "email-sender/internal/handlers_api"

	"email-sender/internal/repositories"
	"email-sender/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var err error
	handlerEmail, err := configService()
	if err != nil {
		log.Fatal("Error config service:", err)
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/process-request-email", handlerEmail.SendEmail).Methods("POST")
	headersOk := handlers.AllowedMethods([]string{"X-Requested-With", "Content-Type", "Authorization", "X-user"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"}) // Ajusta esto para ser más restrictivo si es necesario

	corsRouter := handlers.CORS(originsOk, headersOk, methodsOk)(r)

	// Inicia el servidor HTTP
	port := ":8080"
	fmt.Println("Servidor web iniciado en el puerto", port)
	if err := http.ListenAndServe(port, corsRouter); err != nil {
		panic(err)
	}

}

func configService() (*handlersInternal.EmailHandler, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
		return nil, err
	}
	dialer := clients.NewEmailSender(config.EmailConfig)
	db, err := clients.NewPostgresClient(config.DBConfig.Host,
		config.DBConfig.Port,
		config.DBConfig.Username,
		config.DBConfig.Password,
		config.DBConfig.Database)
	// Check if the connection was successful
	if err != nil {
		log.Println("Error connecting to database:", err)
		return nil, err
	}
	// Initialize repositories
	templateRepo := repositories.NewTemplateRepository(db)
	// Initialize services
	emailService := services.NewEmailService(dialer, templateRepo)
	// Initialize handlers
	emailHandler := handlersInternal.NewEmailHandler(*emailService)
	return emailHandler, nil
}
