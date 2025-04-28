package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"transaction-processor/internal/models"
	"transaction-processor/internal/repositories"
	"transaction-processor/internal/utils"
)

type TransactionService struct {
	transactionRepo *repositories.TransactionRepository
	clientRepo      *repositories.ClientRepository
	emailFrom       string
	bodyTemplate    string
	messageTemplate string
}

func NewTransactionService(clientRepo *repositories.ClientRepository, transactionRepo *repositories.TransactionRepository,
	emailFrom, bodyTemplate, messageTemplate string) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		clientRepo:      clientRepo,
		emailFrom:       emailFrom,
		bodyTemplate:    bodyTemplate,
		messageTemplate: messageTemplate,
	}
}

func (ts *TransactionService) ProcessTransactionFileAndSendEmial(filePath string) error {
	rows := utils.ReadCsvFile(filePath, ",")
	clientId := strings.Split(filePath, ".")[0]
	client, err := ts.getClientDetails(clientId)
	if err != nil {
		return fmt.Errorf("error getting client details: %w", err)
	}

	transactionResume, err := ts.getTransactionResume(rows)
	if err != nil {
		return err
	}
	err = ts.saveTransactions(clientId, transactionResume.Transactions)
	if err != nil {
		return fmt.Errorf("error saving transactions: %w", err)
	}
	subject := fmt.Sprintf("Transaction Resume for %s", client.Name)
	transactionEmail := models.TransactionEmail{
		UserName:          client.Name,
		TransactionResume: transactionResume,
	}
	err = ts.sendEmail(transactionEmail, client.Email, subject, ts.emailFrom,
		ts.bodyTemplate, ts.messageTemplate, []string{})
	if err != nil {
		return fmt.Errorf("error sending emial: %w", err)
	}
	return nil
}

func (ts *TransactionService) getTransactionResume(rows [][]string) (models.TransactionResume, error) {
	creditTransactions := make([]float64, 0)
	debitTransactions := make([]float64, 0)
	monthTransaccions := make(map[string]int, 0)
	allTransactions := make([]models.Transaction, 0)
	totalBalance := 0.0

	for i := range rows {
		if i == 0 {
			continue
		}
		month := utils.GetTransactionMonth(rows[i][1])
		number, err := strconv.ParseFloat(rows[i][2], 64)
		if err != nil {
			return models.TransactionResume{}, fmt.Errorf("error converting string to float: %w", err)
		}

		monthString := utils.GetMonthByNumber(month)
		monthTransaccions[monthString]++
		transactionType := utils.GetTransactionType(number)

		totalBalance += number

		if transactionType == 1 {
			creditTransactions = append(creditTransactions, number)
		} else if transactionType == 2 {
			debitTransactions = append(debitTransactions, number)
		}

		transaction := models.Transaction{
			Amount: number,
			Type:   transactionType,
			Date:   utils.ParseDate(rows[i][1]),
		}
		allTransactions = append(allTransactions, transaction)
	}

	transactionResume := models.TransactionResume{
		TotalBalance:              totalBalance,
		TotalTransactions:         utils.GetTotalTransactions(creditTransactions) + utils.GetTotalTransactions(debitTransactions),
		AverageCreditTransactions: utils.GetAverageTransaction(creditTransactions),
		AverageDebitTransactions:  utils.GetAverageTransaction(debitTransactions),
		TotalCreditTransactions:   utils.GetTotalTransactions(creditTransactions),
		TotalDebitTransactions:    utils.GetTotalTransactions(debitTransactions),
		TotalTransactionsByMonth:  monthTransaccions,
		Transactions:              allTransactions,
	}

	return transactionResume, nil
}

func (ts *TransactionService) getClientDetails(clientId string) (models.Client, error) {
	client, err := ts.clientRepo.GetClientByID(clientId)
	if err != nil {
		return models.Client{}, fmt.Errorf("error getting client details: %w", err)
	}
	return *client, nil
}

func (ts *TransactionService) saveTransactions(clientId string, transaction []models.Transaction) error {
	for _, t := range transaction {
		err := ts.transactionRepo.SaveTransaction(clientId, t)
		if err != nil {
			return fmt.Errorf("error saving transaction: %w", err)
		}
	}
	return nil
}

func (ts *TransactionService) sendEmail(body any, to string, subject string, from string, bodyTemplate string,
	messageTemplate string, attachments []string) error {
	genericEmail := models.GenericEmail{
		To:              to,
		Subject:         subject,
		From:            from,
		Body:            body,
		BodyTemplate:    bodyTemplate,
		MessageTemplate: messageTemplate,
		Attachments:     attachments,
	}
	// Aquí conviertes body a JSON
	jsonBody, err := json.MarshalIndent(genericEmail.Body, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling email body: %w", err)
	}

	fmt.Println("Sending email with body:")
	fmt.Println(string(jsonBody)) // Imprime el body como JSON bonito
	return nil
}
