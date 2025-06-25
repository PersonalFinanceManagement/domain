package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// TransactionType indicates if its an transfer / expense.
type TransactionType string

const (
	Expense  TransactionType = "EXPENSE"
	Income   TransactionType = "INCOME"
	Transfer TransactionType = "TRANSFER"
)

type MethodOfPayment string

const (
	OnlineTransfer MethodOfPayment = "ONLINE_TRANSFER"
	Upi            MethodOfPayment = "UPI"
	CashTransfer   MethodOfPayment = "CASH"
	AutoCredit     MethodOfPayment = "AUTO_CREDIT"
	AutoDebit      MethodOfPayment = "AUTO_DEBIT"
)

type Transaction struct {
	ID      string
	Amount  int64
	Pending bool
	Type    TransactionType

	SourceAccountID      string
	DestinationAccountID string
	Payee                string
	CategoryId           string
	Description          string
	MethodOfPayment      MethodOfPayment

	Created time.Time
	Updated time.Time
}

func NewTransaction(amount int64, pending bool, transactionType TransactionType,
	sourceAccountId, destinationAccountId, payee, categoryId, description string, methodOfPayment MethodOfPayment) (*Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount for transaction cannot be 0 or negative")
	}
	if sourceAccountId == "" {
		return nil, errors.New("transaction needs the sourceaccount")
	}
	return &Transaction{
		ID:                   uuid.New().String(),
		Amount:               amount,
		Pending:              pending,
		Type:                 transactionType,
		SourceAccountID:      sourceAccountId,
		DestinationAccountID: destinationAccountId,
		CategoryId:           categoryId,
		Description:          description,
		MethodOfPayment:      methodOfPayment,
		Created:              time.Now().UTC(),
		Updated:              time.Now().UTC(),
	}, nil
}
