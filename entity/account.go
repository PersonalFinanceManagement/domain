package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AccountType string

const (
	CurrentAccount   AccountType = "CURRENT"
	SavingsAccount   AccountType = "SAVINGS"
	SavingsFDAccount AccountType = "SAVINGS_FD"
	SavingsRDAccount AccountType = "SAVINGS_RD"
	Cash             AccountType = "CASH"
	CreditCard       AccountType = "CREDIT_CARD"
	MutualFund       AccountType = "MUTUAL_FUNDS"
	Stocks           AccountType = "STOCKS"
	LoanAccount      AccountType = "LOAN"
)

type Account struct {
	ID          string
	Name        string
	LinkedBank  string
	AccountType AccountType
	Balance     int64
	Created     time.Time
	Updated     time.Time
	Members     []Member
	Metadata    Metadata
}

func getDefaultAccount() *Account {
	return &Account{
		ID:          uuid.New().String(),
		Balance:     0,
		AccountType: CurrentAccount,
		Created:     time.Now().UTC(),
		Updated:     time.Now().UTC(),
	}
}

func NewAccount(name, linkedBank string, accountType AccountType, balance int64) (*Account, error) {
	if name == "" {
		return nil, errors.New("account name cannot be empty")
	}
	if accountType == "" {
		return nil, errors.New("account type cannot be empty")
	}
	newAccount := getDefaultAccount()
	newAccount.Name = name
	newAccount.LinkedBank = linkedBank
	newAccount.AccountType = accountType
	newAccount.Balance = balance
	return newAccount, nil
}

func (ac *Account) AddMember(members []Member) {
	ac.Members = append(ac.Members, members...)
}

func (ac *Account) AddMetadata(metadata Metadata) {
	ac.Metadata = metadata
}

func (ac *Account) Deposit(amount int64) error {
	if amount <= 0 {
		return errors.New("cannot deposit 0 or less amount")
	}
	ac.Balance += amount
	return nil
}

func (ac *Account) Withdraw(amount int64) error {
	if amount <= 0 {
		return errors.New("cannot withdraw 0 or less amount")
	}
	ac.Balance -= amount
	return nil
}

func (ac *Account) NegativeBalance() bool {
	return ac.Balance <= 0
}
