package models

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
)

type Account struct {
	AccountNumber types.AccountNumber `json:"account_number"`
	CurrencyCode  types.Currency      `json:"currency_code"`
	OwnerName     string              `json:"owner_name"`
	AccountType   types.AccountType   `json:"account_type"`
	Balance       decimal.Decimal     `json:"balance"`
}

type AccountDTO struct {
	AccountNumber types.AccountNumber `json:"account_number"`
	CurrencyCode  types.Currency      `json:"currency_code"`
	OwnerName     string              `json:"owner_name"`
	AccountType   types.AccountType   `json:"account_type"`
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
