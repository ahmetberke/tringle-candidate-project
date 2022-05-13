package cache

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransactionCache(t *testing.T) {
	cache := NewTransactionCache()
	if cache == nil {
		t.Errorf("returned nil, must return TransactionCache")
	}
}

func TestTransactionHistoryCache_AddAccount(t *testing.T) {
	cache := NewTransactionCache()

	t.Run("Success", func(t *testing.T) {
		accountNumber := 1
		err := cache.AddAccount(accountNumber)
		assert.NoError(t, err)
		transactionHistory, err := cache.GetAll(accountNumber)
		assert.NoError(t, err)
		assert.Equal(t, len(transactionHistory), 0)
	})

	t.Run("AlreadyExists", func(t *testing.T) {
		accountNumber := 1
		err := cache.AddAccount(accountNumber)
		assert.NoError(t, err)
		err = cache.AddAccount(accountNumber)
		assert.Error(t, err)
	})

}

func TestTransactionHistoryCache_Create(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		cache := NewTransactionCache()
		accountNumber := 1
		err := cache.AddAccount(accountNumber)
		assert.NoError(t, err)

		transaction := &models.Transaction{
			AccountNumber:   accountNumber,
			Amount:          123,
			TransactionType: "payment",
		}

		transactionR := cache.Create(transaction)
		assert.Equal(t, transactionR, transaction)

		transactionHistory, err := cache.GetAll(accountNumber)
		assert.NoError(t, err)
		assert.Equal(t, len(transactionHistory), 1)

	})

	t.Run("InvalidAccountNumber", func(t *testing.T) {
		cache := NewTransactionCache()
		_, err := cache.GetAll(1)
		assert.Error(t, err)
	})

}
