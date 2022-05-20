package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	AccountNumber   types.AccountNumber
	Amount          decimal.Decimal
	TransactionType types.TransactionType
	CreatedAt       time.Time
}

type TransactionDTO struct {
	AccountNumber   types.AccountNumber   `json:"accountNumber"`
	Amount          float64               `json:"amount"`
	TransactionType types.TransactionType `json:"transactionType"`
	CreatedAt       time.Time             `json:"createdAt"`
}

func (t *Transaction) DTO() *TransactionDTO {

	amountF, _ := t.Amount.Truncate(2).Float64()

	return &TransactionDTO{
		AccountNumber:   t.AccountNumber,
		Amount:          amountF,
		TransactionType: t.TransactionType,
		CreatedAt:       t.CreatedAt,
	}
}

func (td *TransactionDTO) Normal() *Transaction {
	return &Transaction{
		AccountNumber:   td.AccountNumber,
		Amount:          decimal.NewFromFloat(td.Amount),
		TransactionType: td.TransactionType,
		CreatedAt:       td.CreatedAt,
	}
}
