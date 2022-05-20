package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
)

type Account struct {
	AccountNumber types.AccountNumber
	CurrencyCode  types.Currency
	OwnerName     string
	AccountType   types.AccountType
	Balance       decimal.Decimal
}

type AccountDTO struct {
	AccountNumber types.AccountNumber `json:"accountNumber"`
	CurrencyCode  types.Currency      `json:"currencyCode"`
	OwnerName     string              `json:"ownerName"`
	AccountType   types.AccountType   `json:"accountType"`
	Balance       float64             `json:"balance"`
}

func (account *Account) DTO() *AccountDTO {

	balanceF, _ := account.Balance.Truncate(2).Float64()

	return &AccountDTO{
		AccountNumber: account.AccountNumber,
		CurrencyCode:  account.CurrencyCode,
		OwnerName:     account.OwnerName,
		AccountType:   account.AccountType,
		Balance:       balanceF,
	}
}

func (accountDTO *AccountDTO) Normal() *Account {
	return &Account{
		AccountNumber: accountDTO.AccountNumber,
		CurrencyCode:  accountDTO.CurrencyCode,
		OwnerName:     accountDTO.OwnerName,
		AccountType:   accountDTO.AccountType,
		Balance:       decimal.NewFromFloat(accountDTO.Balance),
	}
}
