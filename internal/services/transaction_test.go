package services

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockATransactionCache struct {
	CreateMock     func(transactionHistory *models.Transaction) *models.Transaction
	AddAccountMock func(accountNumber types.AccountNumber) error
	GetAllMock     func(accountNumber types.AccountNumber) ([]*models.Transaction, error)
}

func (m *mockATransactionCache) Create(transactionHistory *models.Transaction) *models.Transaction {
	return m.CreateMock(transactionHistory)
}

func (m *mockATransactionCache) AddAccount(accountNumber types.AccountNumber) error {
	return m.AddAccountMock(accountNumber)
}

func (m *mockATransactionCache) GetAll(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
	return m.GetAllMock(accountNumber)
}

func TestTransactionService_NewPayment(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				var account *models.Account
				if accountNumber == 1 {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Ahmet Berke",
						AccountType:   types.Individual,
						Balance:       decimal.NewFromFloat(float64(500)),
					}
				} else {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Apple",
						AccountType:   types.Corporate,
						Balance:       decimal.NewFromFloat(float64(100)),
					}
				}
				return account, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		payment := &models.Payment{
			SenderAccount:   types.AccountNumber(1),
			ReceiverAccount: types.AccountNumber(2),
			Amount:          decimal.NewFromFloat(50),
		}

		_, err := transactionService.NewPayment(payment)
		assert.NoError(t, err)

	})
	t.Run("InvalidAccountTypeForSender", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				var account *models.Account
				if accountNumber == 1 {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Ahmet Berke",
						AccountType:   types.Corporate,
						Balance:       decimal.NewFromFloat(float64(500)),
					}
				} else {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Apple",
						AccountType:   types.Corporate,
						Balance:       decimal.NewFromFloat(float64(100)),
					}
				}
				return account, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		payment := &models.Payment{
			SenderAccount:   types.AccountNumber(1),
			ReceiverAccount: types.AccountNumber(2),
			Amount:          decimal.NewFromFloat(50),
		}

		_, err := transactionService.NewPayment(payment)
		assert.Error(t, err)

	})
	t.Run("InvalidAccountTypeForReceiver", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				var account *models.Account
				if accountNumber == 1 {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Ahmet Berke",
						AccountType:   types.Individual,
						Balance:       decimal.NewFromFloat(float64(500)),
					}
				} else {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Apple S",
						AccountType:   types.Individual,
						Balance:       decimal.NewFromFloat(float64(100)),
					}
				}
				return account, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		payment := &models.Payment{
			SenderAccount:   types.AccountNumber(1),
			ReceiverAccount: types.AccountNumber(2),
			Amount:          decimal.NewFromFloat(50),
		}

		_, err := transactionService.NewPayment(payment)
		assert.Error(t, err)

	})
	t.Run("InvalidAccountTypeForBoth", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				var account *models.Account
				if accountNumber == 1 {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Ahmet Berke",
						AccountType:   types.Corporate,
						Balance:       decimal.NewFromFloat(float64(500)),
					}
				} else {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Apple S",
						AccountType:   types.Individual,
						Balance:       decimal.NewFromFloat(float64(100)),
					}
				}
				return account, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		payment := &models.Payment{
			SenderAccount:   types.AccountNumber(1),
			ReceiverAccount: types.AccountNumber(2),
			Amount:          decimal.NewFromFloat(50),
		}

		_, err := transactionService.NewPayment(payment)
		assert.Error(t, err)

	})
	t.Run("InvalidCurrencyCode", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				var account *models.Account
				if accountNumber == 1 {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.USD,
						OwnerName:     "Ahmet Berke",
						AccountType:   types.Corporate,
						Balance:       decimal.NewFromFloat(float64(500)),
					}
				} else {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Apple",
						AccountType:   types.Individual,
						Balance:       decimal.NewFromFloat(float64(100)),
					}
				}
				return account, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		payment := &models.Payment{
			SenderAccount:   types.AccountNumber(1),
			ReceiverAccount: types.AccountNumber(2),
			Amount:          decimal.NewFromFloat(50),
		}

		_, err := transactionService.NewPayment(payment)
		assert.Error(t, err)

	})
	t.Run("InsufficientBalance", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				var account *models.Account
				if accountNumber == 1 {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Ahmet Berke",
						AccountType:   types.Individual,
						Balance:       decimal.NewFromFloat(float64(10)),
					}
				} else {
					account = &models.Account{
						AccountNumber: accountNumber,
						CurrencyCode:  types.TRY,
						OwnerName:     "Apple",
						AccountType:   types.Corporate,
						Balance:       decimal.NewFromFloat(float64(100)),
					}
				}
				return account, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		payment := &models.Payment{
			SenderAccount:   types.AccountNumber(1),
			ReceiverAccount: types.AccountNumber(2),
			Amount:          decimal.NewFromFloat(50),
		}

		_, err := transactionService.NewPayment(payment)
		assert.Error(t, err)

	})
}

func TestTransactionService_NewDeposit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return &models.Account{
					AccountNumber: accountNumber,
					CurrencyCode:  types.TRY,
					OwnerName:     "Ahmet Berke",
					AccountType:   types.Individual,
					Balance:       decimal.NewFromFloat(float64(500)),
				}, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		deposit := &models.Deposit{
			AccountNumber: types.AccountNumber(1),
			Amount:        decimal.NewFromFloat(200),
		}

		transaction, err := transactionService.NewDeposit(deposit)
		assert.NoError(t, err)

		assert.True(t, deposit.Amount.Equal(transaction.Amount))

	})
	t.Run("InvalidAccountType", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return &models.Account{
					AccountNumber: accountNumber,
					CurrencyCode:  types.TRY,
					OwnerName:     "Ahmet Berke",
					AccountType:   types.Corporate,
					Balance:       decimal.NewFromFloat(float64(500)),
				}, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		deposit := &models.Deposit{
			AccountNumber: types.AccountNumber(1),
			Amount:        decimal.NewFromFloat(200),
		}

		_, err := transactionService.NewDeposit(deposit)
		assert.Error(t, err)
	})
}

func TestTransactionService_NewWithdraw(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return &models.Account{
					AccountNumber: accountNumber,
					CurrencyCode:  types.TRY,
					OwnerName:     "Ahmet Berke",
					AccountType:   types.Individual,
					Balance:       decimal.NewFromFloat(float64(500)),
				}, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		withdraw := &models.Withdraw{
			AccountNumber: types.AccountNumber(1),
			Amount:        decimal.NewFromFloat(200),
		}

		transaction, err := transactionService.NewWithdraw(withdraw)
		assert.NoError(t, err)

		assert.True(t, withdraw.Amount.Equal(transaction.Amount))

	})
	t.Run("InvalidAccountType", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return &models.Account{
					AccountNumber: accountNumber,
					CurrencyCode:  types.TRY,
					OwnerName:     "Ahmet Berke",
					AccountType:   types.Corporate,
					Balance:       decimal.NewFromFloat(float64(500)),
				}, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		withdraw := &models.Withdraw{
			AccountNumber: types.AccountNumber(1),
			Amount:        decimal.NewFromFloat(200),
		}

		_, err := transactionService.NewWithdraw(withdraw)
		assert.Error(t, err)
	})
	t.Run("InsufficientBalance", func(t *testing.T) {
		mockAccountCach := mockAccountCache{
			GetMock: func(accountNumber types.AccountNumber) (*models.Account, error) {
				return &models.Account{
					AccountNumber: accountNumber,
					CurrencyCode:  types.TRY,
					OwnerName:     "Ahmet Berke",
					AccountType:   types.Corporate,
					Balance:       decimal.NewFromFloat(float64(10)),
				}, nil
			},
			UpdateBalanceMock: func(accountNumber types.AccountNumber, balance decimal.Decimal) error {
				return nil
			},
		}
		mockTransactionCach := mockATransactionCache{
			GetAllMock: func(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
				return nil, errors.New("this account has no transaction history")
			},
			AddAccountMock: func(accountNumber types.AccountNumber) error {
				return nil
			},
			CreateMock: func(transactionHistory *models.Transaction) *models.Transaction {
				return transactionHistory
			},
		}
		transactionService := NewTransactionService(&mockAccountCach, &mockTransactionCach)

		withdraw := &models.Withdraw{
			AccountNumber: types.AccountNumber(1),
			Amount:        decimal.NewFromFloat(200),
		}

		_, err := transactionService.NewWithdraw(withdraw)
		assert.Error(t, err)
	})
}
