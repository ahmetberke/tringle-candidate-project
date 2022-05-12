package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
)

type Account struct {
	AccountNumber int               `json:"account_number"`
	CurrencyCode  types.Currency    `json:"currency_code"`
	OwnerName     string            `json:"owner_name"`
	AccountType   types.AccountType `json:"account_type"`
}
