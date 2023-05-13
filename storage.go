package main

import "database/sql"

type storage interface {
	Create(*Account) error
	Delete(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostGresStore struct {
	db *sql.DB
}

func NewPostGresStore() (*PostGresStore, error) {}
