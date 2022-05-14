package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
)

type Transaction struct {
	AccountNumber   types.AccountNumber   `json:"account_number"`
	Amount          float64               `json:"amount"`
	TransactionType types.TransactionType `json:"transaction_type"`
	CreatedAt       int64                 `json:"created_at"`
}
