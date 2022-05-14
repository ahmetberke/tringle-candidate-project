package cache

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"math"
	"sync"
	"time"
)

type AccountCache struct {
	wg       sync.WaitGroup
	mu       sync.Mutex
	accounts map[types.AccountNumber]*models.Account
}

func NewAccountCache() *AccountCache {
	return &AccountCache{
		wg:       sync.WaitGroup{},
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
	account.AccountNumber = types.AccountNumber(time.Now().UnixNano())
	a.accounts[account.AccountNumber] = account
	return account
}

func (a *AccountCache) Delete(accountNumber types.AccountNumber) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.accounts, accountNumber)
}

func (a *AccountCache) UpdateBalance(accountNumber types.AccountNumber, balance float64) error {
	// Locks with mutex to prevent errors from concurrent access
	a.mu.Lock()
	defer a.mu.Unlock()
	account, err := a.Get(accountNumber)
	if err != nil {
		return err
	}
	account.Balance = math.Round(balance*100) / 100
	return err
}
