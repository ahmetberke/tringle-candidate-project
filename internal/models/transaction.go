package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	AccountNumber   types.AccountNumber   `json:"account_number"`
	Amount          decimal.Decimal       `json:"amount"`
	TransactionType types.TransactionType `json:"transaction_type"`
	CreatedAt       time.Time             `json:"created_at"`
}

type TransactionDTO struct {
	AccountNumber   types.AccountNumber   `json:"account_number"`
	Amount          float64               `json:"amount"`
	TransactionType types.TransactionType `json:"transaction_type"`
	CreatedAt       int64                 `json:"created_at"`
}

func (t *Transaction) DTO() *TransactionDTO {

	amountF, _ := t.Amount.Truncate(2).Float64()

	return &TransactionDTO{
		AccountNumber:   t.AccountNumber,
		Amount:          amountF,
		TransactionType: t.TransactionType,
		CreatedAt:       t.CreatedAt.Unix(),
	}
}

func (td *TransactionDTO) Normal() *Transaction {
	return &Transaction{
		AccountNumber:   td.AccountNumber,
		Amount:          decimal.NewFromFloat(td.Amount),
		TransactionType: td.TransactionType,
		CreatedAt:       time.Unix(td.CreatedAt, 0),
	}
}
