package repository

import "a2sv.org/hub/Domain/entity"

type FundRepository interface {
	Create(fund *entity.Fund) error
	GetByID(id uint) (*entity.Fund, error)
	GetAll() ([]*entity.Fund, error)
	Update(fund *entity.Fund) error
	Delete(id uint) error
}
