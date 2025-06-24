package repository

import (
	"github.com/PersonalFinanceManagement/domain/entity"
)

type BudgetRepository interface {
	Create(budget *entity.Budget) error
	GetByID(id string) (*entity.Budget, error)
	Update(budget *entity.Budget) error
	Delete(id string) error
	Clone(newBudgetName string, budget *entity.Budget) (*entity.Budget, error)
}
