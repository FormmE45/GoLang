package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string
	LastName  string
}

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
	CreatedAt time.Time
}

func NewAccount(FirstName string, LastName string) Account {
	return Account{
		FirstName: FirstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now().UTC(),
	}
}
