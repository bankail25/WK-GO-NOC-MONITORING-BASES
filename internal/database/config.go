package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type IDStore struct {
	db *sql.DB
}

func NewIDStore(dbPath string) (*IDStore, error) {

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	createTable := `CREATE TABLE IF NOT EXISTS ids (
		id TEXT PRIMARY KEY,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		return nil, err
	}
	return &IDStore{db: db}, nil

}

func (s *IDStore) CreateID(id string) (bool, error) {
	query := "INSERT OR IGNORE INTO ids (id) VALUES (?)"
	result, err := s.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0, nil // true si se creó, false si ya existía
}

func (s *IDStore) Exists(id string) (bool, error) {
	query := "SELECT 1 FROM ids WHERE id = ? LIMIT 1"
	var exists int
	err := s.db.QueryRow(query, id).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}
