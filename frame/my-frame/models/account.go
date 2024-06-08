package models

// Account 账户信息
type Account struct {
	AccountNumber int    `json:"account_number"`
	Address       string `json:"address"`
	Age           int    `json:"age"`
	Balance       int    `json:"balance"`
	City          string `json:"city"`
	Email         string `json:"email"`
	Employer      string `json:"employer"`
	Firstname     string `json:"firstname"`
	Gender        string `json:"gender"`
	Lastname      string `json:"lastname"`
	State         string `json:"state"`
}

type Address struct {
	Address   string  `json:"address"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	State     string  `json:"state"`
	Street    string  `json:"street"`
	Zip       string  `json:"zip"`
}
