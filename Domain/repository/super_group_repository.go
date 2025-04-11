package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// SuperGroupRepository defines methods for SuperGroup data operations
type SuperGroupRepository interface {
	CreateSuperGroup(SuperGroup *entity.SuperGroup) error

	ListSuperGroup() ([]*entity.SuperGroup, error)

	GetSuperGroupByName(name string) (*entity.SuperGroup, error)
	GetSuperGroupByID(id uint) (*entity.SuperGroup, error)

	UpdateSuperGroup(SuperGroup *entity.SuperGroup) error

	DeleteSuperGroup(id uint) error
}
