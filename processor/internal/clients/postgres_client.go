package clients

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver
)

func NewPostgresClient(host, port, user, password, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	//Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
