package types

type TransactionType string

const (
	Payment  TransactionType = "payment"
	Deposit                  = "deposit"
	Withdraw                 = "withdraw"
)
