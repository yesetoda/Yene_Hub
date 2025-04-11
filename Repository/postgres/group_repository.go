package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"

	"gorm.io/gorm"
)

// GroupRepository implements the repository.GroupRepository interface
type GroupRepository struct {
	db *gorm.DB
}

// NewGroupRepository creates a new group repository instance
func NewGroupRepository(db *gorm.DB) repository.GroupRepository {
	return &GroupRepository{db: db}
}

// Create creates a new group
func (repo *GroupRepository) Create(group *entity.Group) (*entity.Group, error) {
	result := repo.db.Create(group)
	if result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}

// GetByID retrieves a group by ID
func (repo *GroupRepository) GetByID(id uint) (*entity.Group, error) {
	var group entity.Group
	result := repo.db.First(&group, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &group, nil
}

// GetByName retrieves a group by name
func (repo *GroupRepository) GetByName(name string) (*entity.Group, error) {
	var group entity.Group
	result := repo.db.Where("name = ?", name).First(&group)
	if result.Error != nil {
		return nil, result.Error
	}
	return &group, nil
}

// Update updates a group
func (repo *GroupRepository) Update(group *entity.Group) (*entity.Group, error) {
	result := repo.db.Save(group)
	if result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}

// Delete deletes a group
func (repo *GroupRepository) Delete(id uint) error {
	result := repo.db.Delete(&entity.Group{}, id)
	return result.Error
}

// List retrieves all groups
func (repo *GroupRepository) List() ([]*entity.Group, error) {
	var groups []*entity.Group
	result := repo.db.Find(&groups)
	if result.Error != nil {
		return nil, result.Error
	}
	return groups, nil
}

// GetGroupsByCountryID retrieves groups by country ID
func (repo *GroupRepository) GetGroupsByCountryID(countryID uint) ([]*entity.Group, error) {
	var groups []*entity.Group
	result := repo.db.Where("country_id = ?", countryID).Find(&groups)
	if result.Error != nil {
		return nil, result.Error
	}
	return groups, nil
}
