package models

type Withdraw struct {
	AccountNumber int `json:"account_number"`
	Amount        int `json:"amount"`
}
