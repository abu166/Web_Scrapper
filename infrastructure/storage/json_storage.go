package storage

import (
	"encoding/json"
	"os"
	"web_scrapper/domain"
)

// JSONStorage implements the Repository interface for JSON storage
type JSONStorage struct {
	filename string
}

// NewJSONStorage creates a new JSONStorage
func NewJSONStorage(filename string) *JSONStorage {
	return &JSONStorage{filename: filename}
}

// SaveTablets saves the tablets to a JSON file
func (s *JSONStorage) SaveTablets(tablets []domain.Tablet) error {
	data, err := json.MarshalIndent(tablets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filename, data, 0644)
}