package repositories

import (
	"database/sql"
	"fmt"
	"transaction-processor/internal/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) SaveTransaction(clientID string, transaction models.Transaction) error {
	query := `
		INSERT INTO transactions (client_id, date, amount, transaction_type_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(query, clientID, transaction.Date, transaction.Amount, transaction.Type)
	if err != nil {
		return fmt.Errorf("error inserting transaction: %w", err)
	}
	return nil
}
