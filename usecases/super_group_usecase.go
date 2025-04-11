package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// SuperGroupUseCase defines methods for super group business logic
type SuperGroupUseCaseInterface interface {
	Create(SuperGroup *entity.SuperGroup) (*entity.SuperGroup, error)
	GetByID(id uint) (*entity.SuperGroup, error)
	GetByName(name string) (*entity.SuperGroup, error)
	Update(SuperGroup *entity.SuperGroup) (error)
	Delete(id uint) error
	List() ([]*entity.SuperGroup, error)
}

// SuperGroupUseCase implements SuperGroupUseCase
type SuperGroupUseCase struct {
	superGroupRepo repository.SuperGroupRepository
}

// NewSuperGroupUseCase creates a new SuperGroupUseCase instance
func NewSuperGroupUseCase(superGroupRepo repository.SuperGroupRepository) *SuperGroupUseCase {
	return &SuperGroupUseCase{
		superGroupRepo: superGroupRepo,
	}
}

// Create creates a new super group
func (u *SuperGroupUseCase) Create(SuperGroup *entity.SuperGroup) (*entity.SuperGroup, error) {
	return SuperGroup ,u.superGroupRepo.CreateSuperGroup(SuperGroup)
}

// GetByID retrieves a super group by ID
func (u *SuperGroupUseCase) GetByID(id uint) (*entity.SuperGroup, error) {
	return u.superGroupRepo.GetSuperGroupByID(id)
}

// GetByName retrieves a super group by name
func (u *SuperGroupUseCase) GetByName(name string) (*entity.SuperGroup, error) {
	return u.superGroupRepo.GetSuperGroupByName(name)
}



// Update updates a super group
func (u *SuperGroupUseCase) Update(SuperGroup *entity.SuperGroup) ( error) {
	return u.superGroupRepo.UpdateSuperGroup(SuperGroup)
}

// Delete deletes a super group
func (u *SuperGroupUseCase) Delete(id uint) error {
	return u.superGroupRepo.DeleteSuperGroup(id)
}

// List retrieves all super groups
func (u *SuperGroupUseCase) List() ([]*entity.SuperGroup, error) {
	return u.superGroupRepo.ListSuperGroup()
}
