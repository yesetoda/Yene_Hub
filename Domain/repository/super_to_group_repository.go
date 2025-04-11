package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// SuperToGroupRepository defines methods for SuperToGroup data operations
type SuperToGroupRepository interface {
	CreateSuperToGroup(SuperToGroup *entity.SuperToGroup) error

	ListSuperToGroup() ([]*entity.SuperToGroup, error)

	GetSuperToGroupByID(id uint) (*entity.SuperToGroup, error)

	UpdateSuperToGroup(SuperToGroup *entity.SuperToGroup) error

	DeleteSuperToGroup(id uint) error
}
