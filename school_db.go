package db

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type SchoolDb struct {
	db *sql.DB
}

func InitSchoolDb(dsn string) (*SchoolDb, error) { // name "dsn" -- data source name
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &SchoolDb{db: db}, nil
}

func (repo *SchoolDb) Close() error {
	return repo.db.Close()
}
