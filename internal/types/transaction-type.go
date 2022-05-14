package types

type TransactionType string

const (
	Payment  TransactionType = "payment"
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)
