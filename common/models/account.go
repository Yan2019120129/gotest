package models

import "github.com/brianvoe/gofakeit/v6"

// Account 账户
type Account struct {
	AccountNumber int                   `json:"account_number"`
	Address       *gofakeit.AddressInfo `json:"address"`
	Age           int                   `json:"age"`
	Balance       int                   `json:"balance"`
	City          string                `json:"city"`
	Email         string                `json:"email"`
	Employer      string                `json:"employer"`
	Firstname     string                `json:"firstname"`
	Gender        string                `json:"gender"`
	Lastname      string                `json:"lastname"`
	State         string                `json:"state"`
}
