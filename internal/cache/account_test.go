package cache

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"testing"
)

func TestAccountCache_Create(t *testing.T) {
	accountCache := NewAccountCache()
	eAccount := &models.Account{
		CurrencyCode: types.TRY,
		OwnerName:    "Orkun Demirdağ",
		AccountType:  types.Individual,
	}

	eAccount = accountCache.Create(eAccount)

	expectedANumber := 1
	incomingANumber := eAccount.AccountNumber

	if incomingANumber != expectedANumber {
		t.Errorf("account number is %d but must be %d", incomingANumber, expectedANumber)
	}

	expectedAccount := eAccount
	incomingAccount, err := accountCache.Get(eAccount.AccountNumber)
	if err != nil {
		t.Errorf("account not found")
	}
	if expectedAccount != incomingAccount {
		t.Errorf("account is %p but must be %p", incomingAccount, expectedAccount)
	}

}

func TestAccountCache_Get(t *testing.T) {
	accountCache := NewAccountCache()
	eAccount := &models.Account{
		CurrencyCode: types.TRY,
		OwnerName:    "Orkun Demirdağ",
		AccountType:  types.Individual,
	}
	eAccount = accountCache.Create(eAccount)

	expectedAccount := eAccount
	incomingAccount, err := accountCache.Get(eAccount.AccountNumber)
	if err != nil {
		t.Errorf("account not found")
	}
	if expectedAccount != incomingAccount {
		t.Errorf("account is %p but must be %p", incomingAccount, expectedAccount)
	}
}

func TestAccountCache_Delete(t *testing.T) {
	accountCache := NewAccountCache()
	eAccount := &models.Account{
		CurrencyCode: types.TRY,
		OwnerName:    "Orkun Demirdağ",
		AccountType:  types.Individual,
	}
	eAccount = accountCache.Create(eAccount)
	accountCache.Delete(eAccount.AccountNumber)

	_, err := accountCache.Get(eAccount.AccountNumber)
	if err == nil {
		t.Errorf("account not deleted")
	}

}