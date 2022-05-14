package types

type AccountNumber int64

type AccountType string

const (
	Individual AccountType = "individual"
	Corporate  AccountType = "corporate"
)

type Currency string

const (
	USD Currency = "USD"
	TRY Currency = "TRY"
	EUR Currency = "EUR"
)

type TransactionType string

const (
	Payment  TransactionType = "payment"
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)
