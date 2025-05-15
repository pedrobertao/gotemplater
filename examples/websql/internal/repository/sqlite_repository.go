package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteRepository struct {
	DB *sql.DB
}

func NewSqliteRepository(dbPath string) (*SqliteRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &SqliteRepository{DB: db}, nil
}
