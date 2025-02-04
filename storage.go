package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := os.Getenv("DB")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CretaeAccountTable()
}

func (s *PostgresStore) CretaeAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS Accounts (
    id          SERIAL PRIMARY KEY,
    first_name  VARCHAR(50),
    last_name   VARCHAR(50),
    number      SERIAL,
    balance     SERIAL,
    updated_at  TIMESTAMP,
    created_at  TIMESTAMP
);
`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	query := `
	INSERT INTO Accounts 
	(first_name, last_name, number, balance,updated_at, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id;
	`

	err := s.db.QueryRow(query,
		a.FirstName,
		a.LastName, a.Number,
		a.Balance,
		a.Updated_at,
		a.Created_at,
	).Scan(&a.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateAccount(a *Account, id int) (*Account, error) {
	query := `
	UPDATE Accounts
  SET first_name = $1,
    last_name = $2,
    number = $3,
    balance = $4
  WHERE id = $5
  RETURNING *;
	`

	err := s.db.QueryRow(query,
		a.FirstName,
		a.LastName,
		a.Number,
		a.Balance,
		id,
	).Scan(
		&a.ID,
		&a.FirstName,
		&a.LastName,
		&a.Number,
		&a.Balance,
		&a.Updated_at,
		&a.Created_at,
	)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := `
	DELETE FROM Accounts WHERE id = $1;
	`
	if _, err := s.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query := `
	SELECT * FROM Accounts WHERE id=$1
	`
	account := &Account{}

	err := s.db.QueryRow(query, id).Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.Updated_at,
		&account.Created_at,
	)
	if err != nil {
		return nil, err
	}

	return account, nil
}
