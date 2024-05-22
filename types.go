package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Number     uint64    `json:"number"`
	Balance    uint64    `json:"balance"`
	Updated_at time.Time `json:"updated_at"`
	Created_at time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName:  firstName,
		LastName:   lastName,
		Number:     uint64(rand.Intn(10000000)),
		Updated_at: time.Now().UTC(),
		Created_at: time.Now().UTC(),
	}
}

func UpdatedAccount() *Account {
	return &Account{
		Updated_at: time.Now().UTC(),
	}
}
