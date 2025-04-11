package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// this is vague?
// RoleRepository defines methods for Role data operations
type RoleRepository interface {
	CreateRole(Role *entity.Role) error

	ListRole() ([]*entity.Role, error)

	GetRoleByID(id uint) (*entity.Role, error)
	GetRoleByType(roleType string) ([]*entity.Role, error)

	UpdateRole(Role *entity.Role) error

	DeleteRole(id uint) error
}
