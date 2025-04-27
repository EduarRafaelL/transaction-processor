package models

import "time"

type TransactionResume struct {
	TotalBalance              float64        `json:"total_balance"`
	TotalTransactions         int            `json:"total_transactions"`
	AverageCreditTransactions float64        `json:"average_credit_transactions"`
	AverageDebitTransactions  float64        `json:"average_debit_transactions"`
	TotalCreditTransactions   int            `json:"total_credit_transactions"`
	TotalDebitTransactions    int            `json:"total_debit_transactions"`
	Transactions              []Transaction  `json:"-"`
	TotalTransactionsByMonth  map[string]int `json:"total_transactions_by_month"`
}

type Transaction struct {
	Amount float64   `json:"amount"`
	Type   int       `json:"type"`
	Date   time.Time `json:"date"`
}
