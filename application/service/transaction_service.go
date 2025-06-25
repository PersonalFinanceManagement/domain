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

func (ts *TransactionService) CreateExpense(input CreateTransactionInput) (*entity.Transaction, error) {

	switch input.TransactionType {
	case entity.Expense:
		if input.Payee == "" {
			input.Payee = ts.getDefaultExpensePayee()
		}
		if input.CategoryId == "" {
			input.CategoryId = ts.getDefaultExpenseCategoryID()
		}
		if input.DestinationAccountId == "" {
			input.DestinationAccountId = ts.getDefaultExpenseDestinationAccountId()
		}
	case entity.Income:
		if input.Payee == "" {
			input.Payee = ts.getDefaultIncomePayee()
		}
		if input.CategoryId == "" {
			input.CategoryId = ts.getDefaultIncomeCategoryID()
		}
		if input.SourceAccountId == "" {
			input.SourceAccountId = ts.getDefaultIncomeSourceAccountId()
		}
	case entity.Transfer:
		if input.Payee == "" {
			input.Payee = ts.getDefaultTransferPayee()
		}
		if input.CategoryId == "" {
			input.CategoryId = ts.getDefaultTransferCategoryID()
		}
		if input.SourceAccountId == "" {
			input.SourceAccountId = ts.getDefaultTransferSourceAccountId()
		}
		if input.DestinationAccountId == "" {
			input.DestinationAccountId = ts.getDefaultTransferDestinationAccountId()
		}
	}

	tx, err := entity.NewTransaction(
		input.Amount, input.Pending, input.TransactionType, input.SourceAccountId, input.DestinationAccountId,
		input.Payee, input.CategoryId, input.Description, input.MethodOfPayment)
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
	// return ts.settingsRepo.GetDefaultExpenseDestinationAccountID()
	return "<<EXPENSE_ACCOUNT_ID>>"
}

// get the default transaction account id , something like EXPENSE_EXIT
func (ts *TransactionService) getDefaultExpensePayee() string {
	// return ts.settingsRepo.GetDefaultExpensePayee()
	return "EXPENSOR"
}

// get the default transaction category id
func (ts *TransactionService) getDefaultExpenseCategoryID() string {
	// return ts.settingsRepo.GetDefaultExpenseCategoryID()
	return "<<EXPENSE_CATEGORY_ID>>"
}

// get the default transaction account id , something like a default ID field of the expense type
func (ts *TransactionService) getDefaultIncomeSourceAccountId() string {
	// return ts.settingsRepo.GetDefaultExpenseDestinationAccountID()
	return "<<INCOME_ACCOUNT_ID>>"
}

// get the default transaction account id , something like EXPENSE_EXIT
func (ts *TransactionService) getDefaultIncomePayee() string {
	// return ts.settingsRepo.GetDefaultExpensePayee()
	return "<<INCOMOR>>"
}

// get the default transaction category id
func (ts *TransactionService) getDefaultIncomeCategoryID() string {
	// return ts.settingsRepo.GetDefaultExpenseCategoryID()
	return "<<INCOME_CATEGORY_ID>>"
}

// get the default transaction account id , something like a default ID field of the expense type
func (ts *TransactionService) getDefaultTransferSourceAccountId() string {
	// return ts.settingsRepo.GetDefaultExpenseDestinationAccountID()
	return "<<TRANSFER_SOURCE_ACCOUNT_ID>>"

}

func (ts *TransactionService) getDefaultTransferDestinationAccountId() string {
	// return ts.settingsRepo.GetDefaultExpenseDestinationAccountID()
	return "<<TRANSFER_DESTINATION_ACCOUNT_ID>>"

}

// get the default transaction account id , something like EXPENSE_EXIT
func (ts *TransactionService) getDefaultTransferPayee() string {
	return "<<TRANSFORMER>>"
}

// get the default transaction category id
func (ts *TransactionService) getDefaultTransferCategoryID() string {
	// return ts.settingsRepo.GetDefaultExpenseCategoryID()
	return "<<TRANSFER_CATEGORY_ID>>"
}
