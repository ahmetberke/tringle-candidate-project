package services

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/cache"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
)

type TransactionService struct {
	accountCache            *cache.AccountCache
	transactionHistoryCache *cache.TransactionHistoryCache
}

func NewTransactionService(accountCache *cache.AccountCache,
	transactionHistoryCache *cache.TransactionHistoryCache) *TransactionService {
	return &TransactionService{
		accountCache:            accountCache,
		transactionHistoryCache: transactionHistoryCache,
	}
}

func (ts *TransactionService) NewPayment(payment *models.Payment) (*models.Transaction, error) {
	_, err := ts.transactionHistoryCache.GetAll(payment.SenderAccount)
	if err != nil {
		_ = ts.transactionHistoryCache.AddAccount(payment.SenderAccount)
	}

	sender, err := ts.accountCache.Get(payment.SenderAccount)
	if err != nil {
		return nil, err
	}

	reiever, err := ts.accountCache.Get(payment.ReceiverAccount)
	if err != nil {
		return nil, err
	}

	if sender.AccountType != types.Individual || reiever.AccountType != types.Corporate {
		return nil, errors.New("sender must be individual and receiver must be corporate")
	}

	if sender.CurrencyCode != reiever.CurrencyCode {
		return nil, errors.New("the currency codes of the accounts are not the same")
	}

	if sender.Balance < payment.Amount {
		return nil, errors.New("insufficient balance")
	}

	err = ts.accountCache.UpdateBalance(sender.AccountNumber, sender.Balance-payment.Amount)
	if err != nil {
		return nil, err
	}

	err = ts.accountCache.UpdateBalance(reiever.AccountNumber, reiever.Balance+payment.Amount)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountNumber:   sender.AccountNumber,
		Amount:          payment.Amount,
		TransactionType: types.Payment,
	}

	return ts.transactionHistoryCache.Create(transaction), nil

}

func (ts *TransactionService) NewDeposit(deposit *models.Deposit) (*models.Transaction, error) {
	_, err := ts.transactionHistoryCache.GetAll(deposit.AccountNumber)
	if err != nil {
		_ = ts.transactionHistoryCache.AddAccount(deposit.AccountNumber)
	}

	account, err := ts.accountCache.Get(deposit.AccountNumber)
	if err != nil {
		return nil, err
	}

	if account.AccountType != types.Individual {
		return nil, errors.New("account must be individual")
	}

	err = ts.accountCache.UpdateBalance(account.AccountNumber, account.Balance+deposit.Amount)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountNumber:   account.AccountNumber,
		Amount:          deposit.Amount,
		TransactionType: types.Deposit,
	}

	return ts.transactionHistoryCache.Create(transaction), nil

}

func (ts *TransactionService) NewWithdraw(withdraw *models.Withdraw) (*models.Transaction, error) {
	_, err := ts.transactionHistoryCache.GetAll(withdraw.AccountNumber)
	if err != nil {
		_ = ts.transactionHistoryCache.AddAccount(withdraw.AccountNumber)
	}

	account, err := ts.accountCache.Get(withdraw.AccountNumber)
	if err != nil {
		return nil, err
	}

	if account.AccountType != types.Individual {
		return nil, errors.New("account must be individual")
	}

	if account.Balance < withdraw.Amount {
		return nil, errors.New("insufficient balance")
	}

	err = ts.accountCache.UpdateBalance(account.AccountNumber, account.Balance-withdraw.Amount)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountNumber:   account.AccountNumber,
		Amount:          withdraw.Amount,
		TransactionType: types.Withdraw,
	}

	return ts.transactionHistoryCache.Create(transaction), nil

}

func (ts *TransactionService) GetTransactionHistory(accountNumber int) ([]*models.Transaction, error) {
	return ts.transactionHistoryCache.GetAll(accountNumber)
}
