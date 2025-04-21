package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type roleRepository struct {
	BaseRepository
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{
		BaseRepository: NewBaseRepository(db, "role"),
	}
}

func (r *roleRepository) CreateRole(role *entity.Role) error {
	err := r.db.Create(role).Error
	if err != nil {
		return err
	}

	// Cache the newly created role
	_ = r.cacheDetail("byid", role, role.ID)
	_ = r.cacheDetail("bytype", role, role.Type)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *roleRepository) GetRoleByID(id uint) (*entity.Role, error) {
	var role entity.Role
	err := r.getCachedDetail("byid", &role, func() error {
		return r.db.First(&role, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetRoleByType(roleType string) ([]*entity.Role, error) {
	var roles []*entity.Role
	err := r.getCachedList("bytype", &roles, func() error {
		return r.db.Where("type = ?", roleType).Find(&roles).Error
	}, roleType)
	
	return roles, err
}

func (r *roleRepository) UpdateRole(role *entity.Role) error {
	err := r.db.Save(role).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", role, role.ID)
	_ = r.cacheDetail("bytype", role, role.Type)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *roleRepository) DeleteRole(id uint) error {
	var role entity.Role
	if err := r.db.First(&role, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&role).Error; err != nil {
		return err
	}

	// Invalidate all caches for this role
	r.invalidateAllCache()
	
	return nil
}

func (r *roleRepository) ListRole() ([]*entity.Role, error) {
	var roles []*entity.Role
	err := r.getCachedList("list", &roles, func() error {
		return r.db.Find(&roles).Error
	})
	
	return roles, err
}