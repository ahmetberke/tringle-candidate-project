package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
)

type Withdraw struct {
	AccountNumber types.AccountNumber
	Amount        decimal.Decimal
}

type WithdrawDTO struct {
	AccountNumber types.AccountNumber `json:"accountNumber"`
	Amount        float64             `json:"amount"`
}

func (w *Withdraw) DTO() *WithdrawDTO {
	amountF, _ := w.Amount.Truncate(2).Float64()

	return &WithdrawDTO{
		AccountNumber: w.AccountNumber,
		Amount:        amountF,
	}
}

func (wd *WithdrawDTO) Normal() *Withdraw {
	return &Withdraw{
		AccountNumber: wd.AccountNumber,
		Amount:        decimal.NewFromFloat(wd.Amount),
	}
}
