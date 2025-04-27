package main

import (
	"log"
	"transaction-processor/internal/clients"
	"transaction-processor/internal/config"
	"transaction-processor/internal/repositories"
	"transaction-processor/internal/services"
)

func main() {
	// This is the main entry point for the processor application.
	// The application is designed to process transaction data from a CSV file,
	// perform calculations, and store the results in a PostgreSQL database.
	var err error
	service, err := configService()
	if err != nil {
		log.Fatal("Error config service:", err)
		return
	}
	err = service.ProcessTransactionFileAndSendEmial("12345678.csv")
	if err != nil {
		log.Fatal("Error processing transaction file:", err)
		return
	}
	log.Println("Transaction file processed and email sent successfully.")

}

func configService() (*services.TransactionService, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
		return nil, err
	}
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
	clientRepo := repositories.NewClientRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	// Initialize services
	transactionService := services.NewTransactionService(clientRepo, transactionRepo, config.EmailConfig.From,
		config.EmailConfig.BodyTemplate, config.EmailConfig.MessageTemplate)
	return transactionService, nil

}
