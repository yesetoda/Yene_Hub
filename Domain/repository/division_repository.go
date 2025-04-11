package repository

import "a2sv.org/hub/Domain/entity"

type DivisionRepository interface {
	Create(division *entity.Division) error
	GetByID(id uint) (*entity.Division, error)
	GetByName(name string) (*entity.Division, error)
	GetAll() ([]*entity.Division, error)
	Update(division *entity.Division) error
	Delete(id uint) error
}


