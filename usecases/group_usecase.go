package usecases

import (
	"time"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// GroupUseCase defines methods for group business logic
type GroupUseCase interface {
	Create(group *entity.Group) (*entity.Group, error)
	GetByID(id uint) (*entity.Group, error)
	GetByName(name string) (*entity.Group, error)
	Update(group *entity.Group) (*entity.Group, error)
	Delete(id uint) error
	List() ([]*entity.Group, error)
	GetGroupsByCountryID(countryID uint) ([]*entity.Group, error)
}

// groupUseCase implements GroupUseCase
type groupUseCase struct {
	groupRepo repository.GroupRepository
}

// NewGroupUseCase creates a new GroupUseCase instance
func NewGroupUseCase(groupRepo repository.GroupRepository) GroupUseCase {
	return &groupUseCase{
		groupRepo: groupRepo,
	}
}

// Create creates a new group
func (u *groupUseCase) Create(group *entity.Group) (*entity.Group, error) {
	// Set timestamps
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	// Create group
	return u.groupRepo.Create(group)
}

// GetByID retrieves a group by ID
func (u *groupUseCase) GetByID(id uint) (*entity.Group, error) {
	return u.groupRepo.GetByID(id)
}

// GetByName retrieves a group by name
func (u *groupUseCase) GetByName(name string) (*entity.Group, error) {
	return u.groupRepo.GetByName(name)
}

// Update updates a group
func (u *groupUseCase) Update(group *entity.Group) (*entity.Group, error) {
	existingGroup, err := u.groupRepo.GetByID(group.ID)
	if err != nil {
		return nil, err
	}

	group.CreatedAt = existingGroup.CreatedAt
	group.UpdatedAt = time.Now()

	return u.groupRepo.Update(group)
}

// Delete deletes a group
func (u *groupUseCase) Delete(id uint) error {
	return u.groupRepo.Delete(id)
}

// List retrieves all groups
func (u *groupUseCase) List() ([]*entity.Group, error) {
	return u.groupRepo.List()
}

// GetGroupsByCountryID retrieves groups by country ID
func (u *groupUseCase) GetGroupsByCountryID(countryID uint) ([]*entity.Group, error) {
	return u.groupRepo.GetGroupsByCountryID(countryID)
}
