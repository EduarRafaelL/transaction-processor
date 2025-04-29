package main

import (
	"log"
	"path/filepath"
	"transaction-processor/internal/clients"
	"transaction-processor/internal/config"
	"transaction-processor/internal/repositories"
	"transaction-processor/internal/services"
	"transaction-processor/internal/utils"

	"github.com/fsnotify/fsnotify"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
		return
	}
	service, err := configService(*cfg)
	if err != nil {
		log.Fatal("Error configuring service:", err)
		return
	}

	// Start file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating watcher:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(cfg.InputPath)
	if err != nil {
		log.Fatal("Error watching input path:", err)
		return
	}

	log.Println("Watching directory:", cfg.InputPath)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Create == fsnotify.Create {
				log.Printf("Detected new file: %s\n", event.Name)
				if filepath.Ext(event.Name) == ".csv" {
					log.Printf("Processing CSV file: %s\n", event.Name)
					err := service.ProcessTransactionFileAndSendEmial(event.Name)
					if err != nil {
						log.Printf("Error processing file %s: %v\n", event.Name, err)
						utils.LogError(filepath.Base(event.Name), err)
					} else {
						log.Printf("Successfully processed and sent email for file: %s\n", event.Name)
					}
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}

func configService(config config.Config) (*services.TransactionService, error) {

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
	transactionService := services.NewTransactionService(clientRepo, transactionRepo,
		config.EmailConfig.BodyTemplate, config.EmailConfig.MessageTemplate,
		config.Delimiter, config.EmailConfig.UrlEmailService)
	return transactionService, nil

}
