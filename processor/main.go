package main

import "transaction-processor/internal/services"

func main() {
	// This is the main entry point for the processor application.
	// The application is designed to process transaction data from a CSV file,
	// perform calculations, and store the results in a PostgreSQL database.
	service := services.NewTransactionService()
	err := service.ProcessTransactionFileAndSendEmial("12345678.csv")
	if err != nil {
		panic(err)
	}

}
