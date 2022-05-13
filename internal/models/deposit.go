package models

type Deposit struct {
	AccountNumber int `json:"account_number"`
	Amount        int `json:"amount"`
}
