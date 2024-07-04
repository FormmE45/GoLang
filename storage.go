package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, err
}

func (s *PostgresStore) Init() error {
	return nil
}
func (s *PostgresStore) createTableAccount() error {
	query := `create table if not exists accounts (
	id serialx primary key,
	first_name varchar(50),
	last_name varchar(50),
	account_number serial,
	balance serial,
	created_at timestamp
	)`
	response, err := s.db.Exec(query)
	// if err!=nil{

	// }
	fmt.Printf("%+v\n", response)
	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	queryAddAccountToTableAccount := `
		insert into
		accounts 
		(first_name,last_name,account_number,balance,created_at)
		values
		($1,$2,$3,$4,$5)
	`
	response, err := s.db.Query(queryAddAccountToTableAccount, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Printf("%+v/n", response)
	return nil
}
func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from accounts")
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
