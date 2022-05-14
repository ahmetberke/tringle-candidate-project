package services

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionService_NewPayment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		accountCache := cache.NewAccountCache()
		transactionCache := cache.NewTransactionCache()
		transactionService := NewTransactionService(accountCache, transactionCache)

		account1 := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Mustafa Kırmızı",
			AccountType:  types.Individual,
			Balance:      100,
		})

		account2 := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Apple",
			AccountType:  types.Corporate,
			Balance:      400,
		})

		payment := &models.Payment{
			SenderAccount:   account1.AccountNumber,
			ReceiverAccount: account2.AccountNumber,
			Amount:          50,
		}

		_, err := transactionService.NewPayment(payment)
		assert.NoError(t, err)

		assert.Equal(t, float64(50), account1.Balance)
		assert.Equal(t, float64(450), account2.Balance)

	})

	t.Run("error on account type", func(t *testing.T) {
		accountCache := cache.NewAccountCache()
		transactionCache := cache.NewTransactionCache()
		transactionService := NewTransactionService(accountCache, transactionCache)

		account1 := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Mustafa Kırmızı",
			AccountType:  types.Corporate,
			Balance:      100,
		})

		account2 := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Apple",
			AccountType:  types.Individual,
			Balance:      400,
		})

		payment := &models.Payment{
			SenderAccount:   account1.AccountNumber,
			ReceiverAccount: account2.AccountNumber,
			Amount:          50,
		}

		_, err := transactionService.NewPayment(payment)
		assert.Error(t, err)

	})

	t.Run("error on insufficient balance", func(t *testing.T) {
		accountCache := cache.NewAccountCache()
		transactionCache := cache.NewTransactionCache()
		transactionService := NewTransactionService(accountCache, transactionCache)

		account1 := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Mustafa Kırmızı",
			AccountType:  types.Individual,
			Balance:      100,
		})

		account2 := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Apple",
			AccountType:  types.Corporate,
			Balance:      400,
		})

		payment := &models.Payment{
			SenderAccount:   account1.AccountNumber,
			ReceiverAccount: account2.AccountNumber,
			Amount:          600,
		}

		_, err := transactionService.NewPayment(payment)
		assert.Error(t, err)

	})

}

func TestTransactionService_NewDeposit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		accountCache := cache.NewAccountCache()
		transactionCache := cache.NewTransactionCache()
		transactionService := NewTransactionService(accountCache, transactionCache)

		account := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Mustafa Kırmızı",
			AccountType:  types.Individual,
			Balance:      100,
		})

		deposit := &models.Deposit{
			AccountNumber: account.AccountNumber,
			Amount:        200,
		}

		_, err := transactionService.NewDeposit(deposit)
		assert.NoError(t, err)

		assert.Equal(t, float64(300), account.Balance)

	})

	t.Run("error on account type", func(t *testing.T) {
		accountCache := cache.NewAccountCache()
		transactionCache := cache.NewTransactionCache()
		transactionService := NewTransactionService(accountCache, transactionCache)

		account := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Mustafa Kırmızı",
			AccountType:  types.Corporate,
			Balance:      100,
		})

		deposit := &models.Deposit{
			AccountNumber: account.AccountNumber,
			Amount:        200,
		}

		_, err := transactionService.NewDeposit(deposit)
		assert.Error(t, err)

	})

}

func TestTransactionService_NewWithdraw(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		accountCache := cache.NewAccountCache()
		transactionCache := cache.NewTransactionCache()
		transactionService := NewTransactionService(accountCache, transactionCache)

		account := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Mustafa Kırmızı",
			AccountType:  types.Individual,
			Balance:      100,
		})

		withdraw := &models.Withdraw{
			AccountNumber: account.AccountNumber,
			Amount:        50,
		}

		_, err := transactionService.NewWithdraw(withdraw)
		assert.NoError(t, err)

		assert.Equal(t, float64(50), account.Balance)

	})

	t.Run("error on account type", func(t *testing.T) {
		accountCache := cache.NewAccountCache()
		transactionCache := cache.NewTransactionCache()
		transactionService := NewTransactionService(accountCache, transactionCache)

		account := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Mustafa Kırmızı",
			AccountType:  types.Corporate,
			Balance:      100,
		})

		withdraw := &models.Withdraw{
			AccountNumber: account.AccountNumber,
			Amount:        50,
		}

		_, err := transactionService.NewWithdraw(withdraw)
		assert.Error(t, err)
	})

	t.Run("error on insufficient balance", func(t *testing.T) {
		accountCache := cache.NewAccountCache()
		transactionCache := cache.NewTransactionCache()
		transactionService := NewTransactionService(accountCache, transactionCache)

		account := transactionService.accountCache.Create(&models.Account{
			CurrencyCode: "USD",
			OwnerName:    "Mustafa Kırmızı",
			AccountType:  types.Individual,
			Balance:      100,
		})

		withdraw := &models.Withdraw{
			AccountNumber: account.AccountNumber,
			Amount:        200,
		}

		_, err := transactionService.NewWithdraw(withdraw)
		assert.Error(t, err)
	})

}
