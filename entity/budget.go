package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID          string
	Name        string
	Description string
	Amount      int64
	Spent       int64
	StartDate   time.Time
	EndDate     time.Time
	Created     time.Time
	Updated     time.Time
}

func NewBudget(name string, amount int64, startDate, endDate time.Time) (*Budget, error) {
	if name == "" {
		return nil, errors.New("budget name cannot be empty")
	}
	if amount < 0 {
		return nil, errors.New("budget cannot have negative amount")
	}
	if endDate.Before(startDate) {
		return nil, errors.New("end date of budget has to be after start date")
	}
	return &Budget{
		ID:        uuid.New().String(),
		Name:      name,
		Amount:    amount,
		Spent:     0,
		StartDate: startDate,
		EndDate:   endDate,
		Created:   time.Now().UTC(),
		Updated:   time.Now().UTC(),
	}, nil
}

// RecordSpending updates the amount spent within the budget.
func (b *Budget) RecordSpending(amount int64) error {
	if amount <= 0 {
		return errors.New(fmt.Sprintf("spending amount must be positive. amount:[%d]", amount))
	}
	b.Spent += amount
	return nil
}

// Remaining amount from the budget allocated.
func (b *Budget) RemainingAmount() int64 {
	return b.Amount - b.Spent
}

// IsExceeding checks if the budget is overspending the initial Amount allocated.
func (b *Budget) IsExceeding() bool {
	return b.Spent > b.Amount
}
