package repository

type SettingsRepository interface {
	// Expense Related default fetch
	GetDefaultExpenseDestinationAccountID() string
	GetDefaultExpensePayee() string
	GetDefaultExpenseCategoryID() string
}
