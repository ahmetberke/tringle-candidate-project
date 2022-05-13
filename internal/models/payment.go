package models

type Payment struct {
	SenderAccount   int `json:"sender_account"`
	ReceiverAccount int `json:"receiver_account"`
	Amount          int `json:"amount"`
}
