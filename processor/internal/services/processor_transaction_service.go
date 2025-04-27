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

func (ts *TransactionService) ProcessTransactionFileAndSendEmial(filePath string) (models.TransactionResume, error) {
	rows := utils.ReadCsvFile(filePath, ",")
	transactionResume, err := ts.getTransactionResume(rows)
	if err != nil {
		return models.TransactionResume{}, err
	}
	//TODO: Send email with transactionResume
	return transactionResume, nil
}

func (ts *TransactionService) getTransactionResume(rows [][]string) (models.TransactionResume, error) {
	creditTransactions := make([]float64, 0)
	debitTransactions := make([]float64, 0)
	monthTransaccions := make(map[string]int, 0)
	totalBalance := 0.0

	for i := range rows {
		if i == 0 {
			continue
		}
		month := utils.GateTransactionMonth(rows[i][1])
		number, err := strconv.ParseFloat(rows[i][2], 64)
		if err != nil {
			return models.TransactionResume{}, fmt.Errorf("error converting string to float: %w", err)
		}
		totalBalance += number
		monthString := utils.GetMonthByNumber(month)
		monthTransaccions[monthString]++
		utils.CheckTransacctionType(number, &creditTransactions, &debitTransactions)
	}

	transactionResume := models.TransactionResume{
		TotalBalance:              totalBalance,
		TotalTransactions:         utils.GetTotalTransactions(creditTransactions) + utils.GetTotalTransactions(debitTransactions),
		AverageCreditTransactions: utils.GetAverageTransaction(creditTransactions),
		AverageDebitTransactions:  utils.GetAverageTransaction(debitTransactions),
		TotalCreditTransactions:   utils.GetTotalTransactions(creditTransactions),
		TotalDebitTransactions:    utils.GetTotalTransactions(debitTransactions),
		TotalTransactionsByMonth:  monthTransaccions,
	}
	return transactionResume, nil
}
