package cache

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"sync"
)

var lastAccountNumber = 0

type AccountCache struct {
	wg       sync.WaitGroup
	mu       sync.Mutex
	accounts map[int]*models.Account
}

func NewAccountCache() *AccountCache {
	return &AccountCache{
		wg:       sync.WaitGroup{},
		mu:       sync.Mutex{},
		accounts: make(map[int]*models.Account),
	}
}

func (a *AccountCache) Get(accountNumber int) (*models.Account, error) {
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
	account.AccountNumber = lastAccountNumber
	a.accounts[lastAccountNumber] = account
	return account
}

func (a *AccountCache) Delete(accountNumber int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.accounts, accountNumber)
}

func (a *AccountCache) UpdateBalance(accountNumber int, balance int) error {
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
