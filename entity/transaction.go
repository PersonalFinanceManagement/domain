package entity

import (
	"time"
)

// TransactionType indicates if its an transfer / expense.
type TransactionType string

const (
	Expense  TransactionType = "EXPENSE"
	Income   TransactionType = "INCOME"
	Transfer TransactionType = "TRANSFER"
)

type Transaction struct {
	ID                   string
	Amount               int64
	Pending              bool
	Type                 TransactionType
	SourceAccountID      string
	DestinationAccountID string
	Created              time.Time
	Updated              time.Time
}
