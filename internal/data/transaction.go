package data

import "time"

type TransactionData struct {
	Id        int
	Amount    int
	From      int
	To        int
	Type      string
	CreatedAt time.Time
}

var Transactions = []TransactionData{}
