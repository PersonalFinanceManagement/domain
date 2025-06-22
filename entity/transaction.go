package entity

import (
	"time"
)

// TransactionType indicates if its an transfer / expense.
type TransactionType struct {
}

type Transaction struct {
	ID          string
	Amount      int64
	Pending     bool
	Type        TransactionType
	Source      Account
	Destination Account
	Created     time.Time
	Updated     time.Time
}
