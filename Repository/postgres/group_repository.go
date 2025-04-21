package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type groupRepository struct {
	BaseRepository
}

func NewGroupRepository(db *gorm.DB) repository.GroupRepository {
	return &groupRepository{
		BaseRepository: NewBaseRepository(db, "group"),
	}
}

// Create creates a new group with caching
func (r *groupRepository) Create(group *entity.Group) (*entity.Group, error) {
	err := r.db.Create(group).Error
	if err != nil {
		return nil, err
	}

	// Cache the newly created group
	_ = r.cacheDetail("byid", group, group.ID)

	// Invalidate list caches
	r.invalidateCache("list", "byname")
	
	return group, nil
}

// GetByID retrieves a group by ID using cache
func (r *groupRepository) GetByID(id uint) (*entity.Group, error) {
	var group entity.Group
	err := r.getCachedDetail("byid", &group, func() error {
		return r.db.First(&group, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// GetByName retrieves a group by name using cache
func (r *groupRepository) GetByName(name string) (*entity.Group, error) {
	var group entity.Group
	err := r.getCachedDetail("byname", &group, func() error {
		return r.db.Where("name = ?", name).First(&group).Error
	}, name)
	
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// List retrieves all groups using cache
func (r *groupRepository) List() ([]*entity.Group, error) {
	var groups []*entity.Group
	err := r.getCachedList("list", &groups, func() error {
		return r.db.Find(&groups).Error
	})
	
	return groups, err
}

// Update updates a group and manages cache
func (r *groupRepository) Update(group *entity.Group) (*entity.Group, error) {
	err := r.db.Save(group).Error
	if err != nil {
		return nil, err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", group, group.ID)
	_ = r.cacheDetail("byname", group, group.Name)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return group, nil
}

// Delete deletes a group and invalidates caches
func (r *groupRepository) Delete(id uint) error {
	var group entity.Group
	if err := r.db.First(&group, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&group).Error; err != nil {
		return err
	}

	// Invalidate all caches for this group
	r.invalidateAllCache()
	
	return nil
}

// GetGroupsByCountryID retrieves groups by country ID using cache
func (r *groupRepository) GetGroupsByCountryID(countryID uint) ([]*entity.Group, error) {
	var groups []*entity.Group
	err := r.getCachedList("bycountryid", &groups, func() error {
		return r.db.Where("country_id = ?", countryID).Find(&groups).Error
	}, countryID)
	
	return groups, err
}
