package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewConnection(connString string) (*DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("failed to connection to database: %v:", err)
		return nil, err
	}

	return &DB{db}, nil
}
