package models

import "github.com/ahmetberke/tringle-candidate-project/internal/types"

type Payment struct {
	SenderAccount   types.AccountNumber `json:"sender_account"`
	ReceiverAccount types.AccountNumber `json:"receiver_account"`
	Amount          float64             `json:"amount"`
}
