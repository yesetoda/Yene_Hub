package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) CreateRole(role *entity.Role) error {
	return r.db.Create(role).Error
}

func (r *RoleRepository) GetRoleByID(id uint) (*entity.Role, error) {
	var role entity.Role
	result := r.db.First(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}
func (r *RoleRepository) GetRoleByType(roleType string) ([]*entity.Role, error) {
	var roles []*entity.Role
	result := r.db.Where("role_type = ?", roleType).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
func (r *RoleRepository) UpdateRole(role *entity.Role) error {
	return r.db.Save(role).Error
}
func (r *RoleRepository) DeleteRole(id uint) error {
	return r.db.Delete(&entity.Role{}, id).Error
}

func (r *RoleRepository) ListRole() ([]*entity.Role, error) {
	var roles []*entity.Role
	result := r.db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}