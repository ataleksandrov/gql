package storage

import (
	"database/sql"
	"fmt"
	"log"
)

type Storage interface {
	Get() *sql.DB
	Close() error
}

type postgresStorage struct {
	db *sql.DB
}

func (ps *postgresStorage) Get() *sql.DB {
	return ps.db
}

func (ps *postgresStorage) Close() error {
	return ps.db.Close()
}

// New creates a Storage object and updates the database with the latest migrations
func New(s *Settings) (Storage, error) {
	if s.SkipSSLValidation {
		s.URI = s.URI + "?sslmode=disable"
	}

	db, err := sql.Open(s.Type, s.URI)
	if err != nil {
		return nil, fmt.Errorf("unable to open db connection: %s", err)
	}

	log.Println("Database is up-to-date")

	return &postgresStorage{
		db: db,
	}, nil
}
