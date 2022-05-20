package cache

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"sync"
	"time"
)

type TransactionCache struct {
	mu           sync.Mutex
	transactions map[types.AccountNumber][]*models.Transaction
}

func NewTransactionCache() *TransactionCache {
	return &TransactionCache{
		mu:           sync.Mutex{},
		transactions: make(map[types.AccountNumber][]*models.Transaction),
	}
}

func (tc *TransactionCache) Create(transactionHistory *models.Transaction) *models.Transaction {
	transactionHistory.CreatedAt = time.Now()

	// Locks with mutex to prevent errors from concurrent access
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.transactions[transactionHistory.AccountNumber] = append(tc.transactions[transactionHistory.AccountNumber], transactionHistory)
	return transactionHistory
}

func (tc *TransactionCache) AddAccount(accountNumber types.AccountNumber) error {
	_, ok := tc.transactions[accountNumber]
	if ok {
		return errors.New("this account already has transaction history")
	}

	// Locks with mutex to prevent errors from concurrent access
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.transactions[accountNumber] = []*models.Transaction{}
	return nil
}

func (tc *TransactionCache) GetAll(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
	accounts, ok := tc.transactions[accountNumber]

	if !ok {
		return nil, errors.New("this account has no transaction history")
	}
	return accounts, nil
}
