package services

import (
	"fmt"
	"strconv"
	"transaction-processor/internal/models"
	"transaction-processor/internal/utils"
)

type TransactionService struct {
	//TransactionRepository TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (ts *TransactionService) ProcessTransactionFileAndSendEmial(filePath string) error {
	rows := utils.ReadCsvFile(filePath, ",")
	transactionResume, err := ts.getTransactionResume(rows)
	if err != nil {
		return err
	}
	fmt.Println("Transaction Resume:", transactionResume.Transactions)
	//TODO: Send email with transactionResume
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

func (ts *TransactionService) getClientDetails() (models.Client, error) {
	//TODO: Implement this method to get client details
	return models.Client{}, nil
}

func (ts *TransactionService) saveTransactions(transaction []models.Transaction) error {
	//TODO: Implement this method to save transactions
	return nil
}

func (ts *TransactionService) sendEmail(body any) error {
	genericEmail := models.GenericEmail{
		From:            "",
		To:              "",
		Subject:         "",
		Body:            body,
		BodyTemplate:    "",
		MessageTemplate: "",
		Attachments:     []string{},
	}
	return ts.sendEmail(genericEmail)
}
