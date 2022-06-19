package services

import (
	"errors"
	"github.com/ahmetberke/tringle-candidate-project/internal/models"
	"github.com/ahmetberke/tringle-candidate-project/internal/types"
	"github.com/shopspring/decimal"
)

type TransactionService struct {
	accountCache     accountCache
	transactionCache transactionCache
}

type transactionCache interface {
	Create(transactionHistory *models.Transaction) *models.Transaction
	AddAccount(accountNumber types.AccountNumber) error
	GetAll(accountNumber types.AccountNumber) ([]*models.Transaction, error)
}

func NewTransactionService(ac accountCache,
	tc transactionCache) *TransactionService {
	return &TransactionService{
		accountCache:     ac,
		transactionCache: tc,
	}
}

func (ts *TransactionService) NewPayment(payment *models.Payment) (*models.Transaction, error) {

	if payment.Amount.LessThan(decimal.NewFromInt(0)) {
		return nil, errors.New("amount must be greater than 0")
	}

	_, err := ts.transactionCache.GetAll(payment.SenderAccount)
	if err != nil {
		_ = ts.transactionCache.AddAccount(payment.SenderAccount)
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

	if sender.Balance.LessThan(payment.Amount) {
		return nil, errors.New("insufficient balance")
	}

	err = ts.accountCache.UpdateBalance(sender.AccountNumber, sender.Balance.Sub(payment.Amount))
	if err != nil {
		return nil, err
	}

	err = ts.accountCache.UpdateBalance(reiever.AccountNumber, reiever.Balance.Add(payment.Amount))
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountNumber:   sender.AccountNumber,
		Amount:          payment.Amount,
		TransactionType: types.Payment,
	}

	return ts.transactionCache.Create(transaction), nil

}

func (ts *TransactionService) NewDeposit(deposit *models.Deposit) (*models.Transaction, error) {

	if deposit.Amount.LessThan(decimal.NewFromInt(0)) {
		return nil, errors.New("amount must be greater than 0")
	}

	_, err := ts.transactionCache.GetAll(deposit.AccountNumber)
	if err != nil {
		_ = ts.transactionCache.AddAccount(deposit.AccountNumber)
	}

	account, err := ts.accountCache.Get(deposit.AccountNumber)
	if err != nil {
		return nil, err
	}

	if account.AccountType != types.Individual {
		return nil, errors.New("account must be individual")
	}

	err = ts.accountCache.UpdateBalance(account.AccountNumber, account.Balance.Add(deposit.Amount))
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountNumber:   account.AccountNumber,
		Amount:          deposit.Amount,
		TransactionType: types.Deposit,
	}

	return ts.transactionCache.Create(transaction), nil

}

func (ts *TransactionService) NewWithdraw(withdraw *models.Withdraw) (*models.Transaction, error) {

	if withdraw.Amount.LessThan(decimal.NewFromInt(0)) {
		return nil, errors.New("amount must be greater than 0")
	}

	_, err := ts.transactionCache.GetAll(withdraw.AccountNumber)
	if err != nil {
		_ = ts.transactionCache.AddAccount(withdraw.AccountNumber)
	}

	account, err := ts.accountCache.Get(withdraw.AccountNumber)
	if err != nil {
		return nil, err
	}

	if account.AccountType != types.Individual {
		return nil, errors.New("account must be individual")
	}

	if account.Balance.LessThan(withdraw.Amount) {
		return nil, errors.New("insufficient balance")
	}

	err = ts.accountCache.UpdateBalance(account.AccountNumber, account.Balance.Sub(withdraw.Amount))
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountNumber:   account.AccountNumber,
		Amount:          withdraw.Amount,
		TransactionType: types.Withdraw,
	}

	return ts.transactionCache.Create(transaction), nil

}

func (ts *TransactionService) GetTransactionHistory(accountNumber types.AccountNumber) ([]*models.Transaction, error) {
	return ts.transactionCache.GetAll(accountNumber)
}
