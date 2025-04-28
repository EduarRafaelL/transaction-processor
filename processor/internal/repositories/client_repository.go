package repositories

import (
	"database/sql"
	"transaction-processor/internal/models"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (r *ClientRepository) GetClientByID(clientID string) (*models.Client, error) {
	query := `SELECT id, name, email FROM clients WHERE id = $1`
	row := r.db.QueryRow(query, clientID)

	client := &models.Client{}
	err := row.Scan(&client.ID, &client.Name, &client.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No client found
		}
		return nil, err // Error scanning row
	}
	return client, nil
}
