package storage_user_mails

import "database/sql"

type Storage struct {
	db *sql.DB
}

func New(
	db *sql.DB,
) *Storage {
	return &Storage{
		db: db,
	}
}
