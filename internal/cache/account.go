package cache

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
	"sync"
)

var lastAccountNumber = 0

type AccountCache struct {
	mu       sync.Mutex
	accounts map[types.AccountNumber]*models.Account
}

func NewAccountCache() *AccountCache {
	return &AccountCache{
		mu:       sync.Mutex{},
		accounts: make(map[types.AccountNumber]*models.Account),
	}
}

func (a *AccountCache) Get(accountNumber types.AccountNumber) (*models.Account, error) {
	account, ok := a.accounts[accountNumber]
	if !ok {
		return nil, errors.New("invalid account number")
	}
	return account, nil
}

func (a *AccountCache) Create(account *models.Account) *models.Account {
	// Locks with mutex to prevent errors from concurrent access
	a.mu.Lock()
	defer a.mu.Unlock()
	lastAccountNumber++
	account.AccountNumber = types.AccountNumber(lastAccountNumber)
	a.accounts[account.AccountNumber] = account
	return account
}

func (a *AccountCache) Delete(accountNumber types.AccountNumber) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.accounts, accountNumber)
}

func (a *AccountCache) UpdateBalance(accountNumber types.AccountNumber, balance decimal.Decimal) error {
	// Locks with mutex to prevent errors from concurrent access
	a.mu.Lock()
	defer a.mu.Unlock()
	account, err := a.Get(accountNumber)
	if err != nil {
		return err
	}
	account.Balance = balance
	return err
}
