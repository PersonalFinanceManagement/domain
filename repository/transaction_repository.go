package repository

import (
	"github.com/PersonalFinanceManagement/domain/entity"
)

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	GetByID(id string) (*entity.Transaction, error)
	Update(transaction *entity.Transaction) error
	Delete(id string) error
	Clone(newTransactionName string, transaction *entity.Transaction) (*entity.Transaction, error)
}
