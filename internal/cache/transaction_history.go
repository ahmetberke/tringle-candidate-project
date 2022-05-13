package cache

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"sync"
	"time"
)

type TransactionHistoryCache struct {
	wg           sync.WaitGroup
	mu           sync.Mutex
	transactions map[int][]*models.Transaction
}

func NewTransactionCache() *TransactionHistoryCache {
	return &TransactionHistoryCache{
		wg:           sync.WaitGroup{},
		mu:           sync.Mutex{},
		transactions: make(map[int][]*models.Transaction),
	}
}

func (tc *TransactionHistoryCache) Create(transactionHistory *models.Transaction) *models.Transaction {
	transactionHistory.CreatedAt = time.Now()
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.transactions[transactionHistory.AccountNumber] = append(tc.transactions[transactionHistory.AccountNumber], transactionHistory)
	return transactionHistory
}

func (tc *TransactionHistoryCache) AddAccount(accountNumber int) error {
	_, ok := tc.transactions[accountNumber]
	if ok {
		return errors.New("this account already has transaction history")
	}
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.transactions[accountNumber] = []*models.Transaction{}
	return nil
}

func (tc *TransactionHistoryCache) GetAll(accountNumber int) ([]*models.Transaction, error) {
	accounts, ok := tc.transactions[accountNumber]
	if !ok {
		return nil, errors.New("this account has no transaction history")
	}
	return accounts, nil
}
