package cache

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountCache_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		accountCache := NewAccountCache()
		account := accountCache.Create(&models.Account{
			CurrencyCode: types.TRY,
			OwnerName:    "Ken Thompson",
			AccountType:  types.Individual,
		})
		assert.Equal(t, types.AccountNumber(1), account.AccountNumber)
	})
}

func TestAccountCache_Get(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		accountCache := NewAccountCache()
		account := accountCache.Create(&models.Account{
			CurrencyCode: types.TRY,
			OwnerName:    "Ken Thompson",
			AccountType:  types.Individual,
		})

		iAccount, err := accountCache.Get(account.AccountNumber)
		assert.NoError(t, err)
		assert.Equal(t, iAccount.AccountNumber, account.AccountNumber)
	})
	t.Run("AccountNotFound", func(t *testing.T) {
		accountCache := NewAccountCache()
		_, err := accountCache.Get(types.AccountNumber(1))
		assert.Error(t, err)
	})
}
func TestAccountCache_Delete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		accountCache := NewAccountCache()
		account := accountCache.Create(&models.Account{
			CurrencyCode: types.TRY,
			OwnerName:    "Ken Thompson",
			AccountType:  types.Individual,
		})

		accountCache.Delete(account.AccountNumber)
		_, err := accountCache.Get(account.AccountNumber)
		assert.Error(t, err)
	})
}

func TestAccountCache_UpdateBalance(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		accountCache := NewAccountCache()
		eAccount := &models.Account{
			CurrencyCode: types.TRY,
			OwnerName:    "Orkun Demirdağ",
			AccountType:  types.Individual,
			Balance:      decimal.NewFromFloat(float64(123)),
		}
		eAccount = accountCache.Create(eAccount)
		err := accountCache.UpdateBalance(eAccount.AccountNumber, decimal.NewFromFloat(200))
		assert.NoError(t, err)

		aAccount, err := accountCache.Get(eAccount.AccountNumber)
		amountF, ok := aAccount.Balance.Float64()
		assert.True(t, ok)
		assert.Equal(t, float64(200), amountF)
	})
}
