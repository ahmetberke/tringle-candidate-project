package models

import "github.com/ahmetberke/tringle-candidate-project/internal/types"

type Deposit struct {
	AccountNumber types.AccountNumber `json:"account_number"`
	Amount        float64             `json:"amount"`
}
