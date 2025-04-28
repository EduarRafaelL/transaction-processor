package repositories

import (
	"database/sql"
	"email-sender/internal/models"
)

type TemplateRepository struct {
	db *sql.DB
}

func NewTemplateRepository(db *sql.DB) *TemplateRepository {
	return &TemplateRepository{
		db: db,
	}
}

func (r *TemplateRepository) GetTemplateByName(templateName string) (models.Template, error) {
	query := `SELECT id, name, subject, body FROM templates WHERE name = $1`
	row := r.db.QueryRow(query, templateName)

	var template models.Template
	err := row.Scan(&template.ID, &template.Name, &template.Subject, &template.HtmlBody)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Template{}, nil // No template found
		}
		return models.Template{}, err // Error scanning row
	}
	return template, nil
}
