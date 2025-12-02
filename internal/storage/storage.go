package storage

import (
	"database/sql"
	"pets-planify/internal"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) Tx() (internal.Tx, error) {
	return s.db.Begin()
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}
