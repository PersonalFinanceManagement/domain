package entity

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// TransactionType indicates if its an transfer / expense.
type TransactionType struct {
}

type Transaction struct {
	ID          string
	Amount      float64
	Pending     bool
	Type        TransactionType
	Source      Account
	Destination Account
	Created     time.Time
	Updated     time.Time
}
