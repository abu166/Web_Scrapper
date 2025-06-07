package storage

import (
	"database/sql"
	"fmt"
	"web_scrapper/domain"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Initialize() error {
	query := `
		CREATE TABLE IF NOT EXISTS tablets (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			price FLOAT NOT NULL,
			description TEXT NOT NULL,
			rating INTEGER NOT NULL,
			image_url TEXT NOT NULL
		)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) SaveTablets(tablets []domain.Tablet) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	stmt, err := tx.Prepare("INSERT INTO tablets (title, price, description, rating, image_url) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	for _, tablet := range tablets {
		_, err := stmt.Exec(tablet.Title, tablet.Price, tablet.Description, tablet.Rating, tablet.ImageURL)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert tablet: %v", err)
		}
	}

	return tx.Commit()
}
