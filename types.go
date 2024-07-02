package main

import "time"

type CreateAccountRequest struct {
	FirstName string
	LastName  string
}

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    string
	Balance   int64
	CreatedAt time.Time
}

var id int = 0
var idadrr = &id

func NewAccount(FirstName string, LastName string, Number string) Account {
	*idadrr += 1
	return Account{
		ID:        *idadrr,
		FirstName: FirstName,
		LastName:  LastName,
		Number:    Number,
		CreatedAt: time.Now().UTC(),
	}
}
