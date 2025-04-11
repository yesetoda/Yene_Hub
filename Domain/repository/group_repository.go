package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// GroupRepository defines methods for group data operations
type GroupRepository interface {
	Create(group *entity.Group) (*entity.Group, error)
	GetByID(id uint) (*entity.Group, error)
	GetByName(name string) (*entity.Group, error)
	Update(group *entity.Group) (*entity.Group, error)
	Delete(id uint) error
	List() ([]*entity.Group, error)
	GetGroupsByCountryID(countryID uint) ([]*entity.Group, error)
}
