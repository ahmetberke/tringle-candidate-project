package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
)

type Payment struct {
	SenderAccount   types.AccountNumber `json:"sender_account"`
	ReceiverAccount types.AccountNumber `json:"receiver_account"`
	Amount          decimal.Decimal     `json:"amount"`
}

type PaymentDTO struct {
	SenderAccount   types.AccountNumber `json:"sender_account"`
	ReceiverAccount types.AccountNumber `json:"receiver_account"`
	Amount          float64             `json:"amount"`
}

func (p *Payment) DTO() *PaymentDTO {
	amountF, _ := p.Amount.Truncate(2).Float64()

	return &PaymentDTO{
		SenderAccount:   p.SenderAccount,
		ReceiverAccount: p.ReceiverAccount,
		Amount:          amountF,
	}
}

func (pd *PaymentDTO) Normal() *Payment {
	return &Payment{
		SenderAccount:   pd.SenderAccount,
		ReceiverAccount: pd.ReceiverAccount,
		Amount:          decimal.NewFromFloat(pd.Amount),
	}
}
