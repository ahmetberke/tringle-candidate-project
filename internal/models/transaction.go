package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"time"
)

type Transaction struct {
	AccountNumber   int                   `json:"account_number"`
	Amount          int                   `json:"amount"`
	TransactionType types.TransactionType `json:"transaction_type"`
	CreatedAt       time.Time             `json:"created_at"`
}
