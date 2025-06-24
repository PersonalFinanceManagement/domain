package repository

import (
	"github.com/PersonalFinanceManagement/domain/entity"
)

type AccountRepository interface {
	Create(account *entity.Account) error
	GetByID(id string) (*entity.Account, error)
	Update(account *entity.Account) error
	Delete(id string) error
	Clone(newAccountName string, account *entity.Account) (*entity.Account, error)
}
