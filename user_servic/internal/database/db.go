package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewConnection(connStr string) (*sql.DB, error) {

	return sql.Open("postgres", connStr)
}
