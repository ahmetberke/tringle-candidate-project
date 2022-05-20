package services

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockAccountCache struct {
	GetMock           func(accountNumber types.AccountNumber) (*models.Account, error)
	CreateMock        func(account *models.Account) *models.Account
	DeleteMock        func(accountNumber types.AccountNumber)
	UpdateBalanceMock func(accountNumber types.AccountNumber, balance decimal.Decimal) error
}

func (m *mockAccountCache) Get(accountNumber types.AccountNumber) (*models.Account, error) {
	return m.GetMock(accountNumber)
}

func (m *mockAccountCache) Create(account *models.Account) *models.Account {
	return m.CreateMock(account)
}

func (m *mockAccountCache) Delete(accountNumber types.AccountNumber) {
	m.DeleteMock(accountNumber)
}

func (m *mockAccountCache) UpdateBalance(accountNumber types.AccountNumber, balance decimal.Decimal) error {
	return m.UpdateBalanceMock(accountNumber, balance)
}

func TestAccountService_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			CreateMock: func(account *models.Account) *models.Account {
				return &models.Account{
					AccountNumber: 1,
					CurrencyCode:  account.CurrencyCode,
					OwnerName:     account.OwnerName,
					AccountType:   account.AccountType,
					Balance:       decimal.Decimal{},
				}
			},
		}
		accounService := NewAccountService(&mockAccountCach)

		account := models.Account{
			CurrencyCode: types.USD,
			OwnerName:    "Robert Griesemer",
			AccountType:  types.Individual,
		}

		createdAccount, err := accounService.Create(&account)
		assert.NoError(t, err)
		assert.Equal(t, account.AccountType, createdAccount.AccountType)
		assert.Equal(t, account.OwnerName, createdAccount.OwnerName)
		assert.Equal(t, account.CurrencyCode, createdAccount.CurrencyCode)

		balanceF, _ := createdAccount.Balance.Float64()
		assert.Equal(t, float64(0), balanceF)

	})
	t.Run("InvalidAccountType", func(t *testing.T) {
		mockAccountCach := mockAccountCache{}
		accounService := NewAccountService(&mockAccountCach)

		account := models.Account{
			CurrencyCode: types.USD,
			OwnerName:    "Robert Griesemer",
			AccountType:  types.AccountType("x"),
		}

		_, err := accounService.Create(&account)
		assert.Error(t, err)
	})
	t.Run("InvalidCurrencyCode", func(t *testing.T) {
		mockAccountCach := mockAccountCache{}
		accounService := NewAccountService(&mockAccountCach)

		account := models.Account{
			CurrencyCode: types.Currency("x"),
			OwnerName:    "Robert Griesemer",
			AccountType:  types.Individual,
		}

		_, err := accounService.Create(&account)
		assert.Error(t, err)
	})
	t.Run("InvalidOwnerNameForIndividualAccount", func(t *testing.T) {
		mockAccountCach := mockAccountCache{}
		accounService := NewAccountService(&mockAccountCach)

		account := models.Account{
			CurrencyCode: types.TRY,
			OwnerName:    "Robert",
			AccountType:  types.Individual,
		}

		_, err := accounService.Create(&account)
		assert.Error(t, err)
	})
}

func TestAccountService_FindByAccountNumber(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return &models.Account{
					AccountNumber: accountNumber,
				}, nil
			},
		}
		accounService := NewAccountService(&mockAccountCach)

		account, err := accounService.FindByAccountNumber(1)
		assert.NoError(t, err)
		assert.Equal(t, types.AccountNumber(1), account.AccountNumber)
	})
	t.Run("AccountNotFound", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return nil, errors.New("account not found")
			},
		}
		accounService := NewAccountService(&mockAccountCach)

		_, err := accounService.FindByAccountNumber(1)
		assert.Error(t, err)
	})
}

func TestAccountService_Delete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			DeleteMock: func(accountNumber types.AccountNumber) {
			},
		}
		accounService := NewAccountService(&mockAccountCach)
		accounService.Delete(1)
	})
}
