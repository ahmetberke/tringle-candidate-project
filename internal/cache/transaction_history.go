package cache

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"math"
	"sync"
	"time"
)

type TransactionHistoryCache struct {
	wg           sync.WaitGroup
	mu           sync.Mutex
	transactions map[types.AccountNumber][]*models.Transaction
}

func NewTransactionCache() *TransactionHistoryCache {
	return &TransactionHistoryCache{
		wg:           sync.WaitGroup{},
		mu:           sync.Mutex{},
		transactions: make(map[types.AccountNumber][]*models.Transaction),
	}
}

func (tc *TransactionHistoryCache) Create(transactionHistory *models.Transaction) *models.Transaction {
	transactionHistory.CreatedAt = time.Now().Unix()

	// Round to 2 decimal
	transactionHistory.Amount = math.Round(transactionHistory.Amount*100) / 100

	// Locks with mutex to prevent errors from concurrent access
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.transactions[transactionHistory.AccountNumber] = append(tc.transactions[transactionHistory.AccountNumber], transactionHistory)
	return transactionHistory
}

func (tc *TransactionHistoryCache) AddAccount(accountNumber types.AccountNumber) error {
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

func (tc *TransactionHistoryCache) GetAll(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
	accounts, ok := tc.transactions[accountNumber]

	if !ok {
		return nil, errors.New("this account has no transaction history")
	}
	return accounts, nil
}
