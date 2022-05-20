package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
)

type Deposit struct {
	AccountNumber types.AccountNumber
	Amount        decimal.Decimal
}

type DepositDTO struct {
	AccountNumber types.AccountNumber `json:"accountNumber"`
	Amount        float64             `json:"amount"`
}

func (d *Deposit) DTO() *DepositDTO {
	amountF, _ := d.Amount.Truncate(2).Float64()

	return &DepositDTO{
		AccountNumber: d.AccountNumber,
		Amount:        amountF,
	}
}

func (dd *DepositDTO) Normal() *Deposit {
	return &Deposit{
		AccountNumber: dd.AccountNumber,
		Amount:        decimal.NewFromFloat(dd.Amount),
	}
}
