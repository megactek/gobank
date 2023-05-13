package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostGresStore struct {
	db *sql.DB
}

func NewPostGresStore() (*PostGresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostGresStore{db: db}, nil
}

func (s *PostGresStore) CreateAccount(*Account) error {
	return nil
}
func (s *PostGresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostGresStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostGresStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
