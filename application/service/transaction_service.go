package service

import (
	"github.com/PersonalFinanceManagement/domain/entity"
	"github.com/PersonalFinanceManagement/domain/repository"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepository
	settingsRepo    repository.SettingsRepository
}

func NewTransactionService(transRepo repository.TransactionRepository, settingsRepo repository.SettingsRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transRepo,
		settingsRepo:    settingsRepo,
	}
}

type CreateTransactionInput struct {
	Amount               int64
	Pending              bool
	TransactionType      entity.TransactionType
	SourceAccountId      string
	DestinationAccountId string
	Payee                string
	CategoryId           string
	Description          string
	MethodOfPayment      entity.MethodOfPayment
}

func (ts *TransactionService) CreateExpense(transInput CreateTransactionInput) (*entity.Transaction, error) {
	if transInput.TransactionType == entity.Expense {
		// get the default transactionAccountId for expense type transaction
		if transInput.DestinationAccountId == "" {
			transInput.DestinationAccountId = ts.getDefaultExpenseDestinationAccountId()
		}
		if transInput.Payee == "" {
			transInput.Payee = ts.getDefaultExpensePayee()
		}
		if transInput.CategoryId == "" {
			transInput.CategoryId = ts.getDefaultExpenseCategoryID()
		}
	}
	tx, err := entity.NewTransaction(
		transInput.Amount, transInput.Pending, transInput.TransactionType, transInput.SourceAccountId, transInput.DestinationAccountId,
		transInput.Payee, transInput.CategoryId, transInput.Description, transInput.MethodOfPayment)
	if err != nil {
		return nil, err
	}
	err = ts.transactionRepo.Create(tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// get the default transaction account id , something like a default ID field of the expense type
func (ts *TransactionService) getDefaultExpenseDestinationAccountId() string {
	return ts.settingsRepo.GetDefaultExpenseDestinationAccountID()
}

// get the default transaction account id , something like EXPENSE_EXIT
func (ts *TransactionService) getDefaultExpensePayee() string {
	return ts.settingsRepo.GetDefaultExpensePayee()
}

// get the default transaction category id
func (ts *TransactionService) getDefaultExpenseCategoryID() string {
	return ts.settingsRepo.GetDefaultExpenseCategoryID()
}
