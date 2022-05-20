package services

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
	"strings"
)

type AccountService struct {
	Cache accountCache
}

type accountCache interface {
	Get(accountNumber types.AccountNumber) (*models.Account, error)
	Create(account *models.Account) *models.Account
	Delete(accountNumber types.AccountNumber)
	UpdateBalance(accountNumber types.AccountNumber, balance decimal.Decimal) error
}

func NewAccountService(cache *cache.AccountCache) *AccountService {
	return &AccountService{Cache: cache}
}

func (as *AccountService) FindByAccountNumber(accountNumber types.AccountNumber) (*models.Account, error) {
	if accountNumber < 0 {
		return nil, errors.New("account number cannot be negative")
	}
	return as.Cache.Get(accountNumber)
}

func (as *AccountService) Create(account *models.Account) (*models.Account, error) {

	// Checking valid account type
	switch account.CurrencyCode {
	case types.TRY, types.EUR, types.USD:
	default:
		return nil, errors.New("invalid currency code")
	}

	// Checking valid account type
	switch account.AccountType {
	case types.Individual:
		// Checking valid owner name for individual accounts
		res := strings.Split(account.OwnerName, " ")
		if len(res) < 2 {
			return nil, errors.New("invalid owner name")
		}

	case types.Corporate:
	default:
		return nil, errors.New("invalid account type")
	}
	return as.Cache.Create(account), nil
}

func (as *AccountService) Delete(accountNumber types.AccountNumber) {
	as.Cache.Delete(accountNumber)
}
