package services

import (
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
)

type AccountService struct {
	Cache *cache.AccountCache
}

func NewAccountService(cache *cache.AccountCache) *AccountService {
	return &AccountService{Cache: cache}
}

func (as *AccountService) FindByAccountNumber(accountNumber int) (*models.Account, error) {
	return as.Cache.Get(accountNumber)
}

func (as *AccountService) Create(account *models.Account) *models.Account {
	return as.Cache.Create(account)
}

func (as *AccountService) Delete(accountNumber int) {
	as.Cache.Delete(accountNumber)
}
