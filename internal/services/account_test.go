package services

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"testing"
)

func TestAccountService_Create(t *testing.T) {
	accountCache := cache.NewAccountCache()
	accountService := NewAccountService(accountCache)

	eAccount := &models.Account{
		CurrencyCode: types.TRY,
		OwnerName:    "Orkun Demirdağ",
		AccountType:  types.Individual,
	}

	eAccount = accountService.Create(eAccount)

	expectedANumber := 1
	incomingANumber := eAccount.AccountNumber

	if incomingANumber != expectedANumber {
		t.Errorf("account number is %d but must be %d", incomingANumber, expectedANumber)
	}

	expectedAccount := eAccount
	incomingAccount, err := accountService.FindByAccountNumber(eAccount.AccountNumber)
	if err != nil {
		t.Errorf("account not found")
	}
	if expectedAccount != incomingAccount {
		t.Errorf("account is %p but must be %p", incomingAccount, expectedAccount)
	}

}

func TestAccountService_FindByAccountNumber(t *testing.T) {
	accountCache := cache.NewAccountCache()
	accountService := NewAccountService(accountCache)

	eAccount := &models.Account{
		CurrencyCode: types.TRY,
		OwnerName:    "Orkun Demirdağ",
		AccountType:  types.Individual,
	}
	eAccount = accountService.Create(eAccount)

	expectedAccount := eAccount
	incomingAccount, err := accountService.FindByAccountNumber(eAccount.AccountNumber)
	if err != nil {
		t.Errorf("account not found")
	}
	if expectedAccount != incomingAccount {
		t.Errorf("account is %p but must be %p", incomingAccount, expectedAccount)
	}
}

func TestAccountService_Delete(t *testing.T) {
	accountCache := cache.NewAccountCache()
	accountService := NewAccountService(accountCache)
	eAccount := &models.Account{
		CurrencyCode: types.TRY,
		OwnerName:    "Orkun Demirdağ",
		AccountType:  types.Individual,
	}
	eAccount = accountService.Create(eAccount)
	accountCache.Delete(eAccount.AccountNumber)

	_, err := accountService.FindByAccountNumber(eAccount.AccountNumber)
	if err == nil {
		t.Errorf("account not deleted")
	}

}
