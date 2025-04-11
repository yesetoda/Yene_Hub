package usecases

import (
	"time"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// RoleUseCase defines methods for role business logic
type RoleUseCaseInterface interface {
	Create(role *entity.Role) (*entity.Role, error)
	GetByID(id uint) (*entity.Role, error)
	GetByType(roleType string) ([]*entity.Role, error)
	Update(role *entity.Role) (*entity.Role, error)
	Delete(id uint) error
	List() ([]*entity.Role, error)
}

// RoleUseCase implements RoleUseCase
type RoleUseCase struct {
	RoleRepository repository.RoleRepository
}

// NewRoleUseCase creates a new RoleUseCase instance
func NewRoleUseCase(RoleRepository repository.RoleRepository) *RoleUseCase {
	return &RoleUseCase{
		RoleRepository: RoleRepository,
	}
}

// Create creates a new role
func (u *RoleUseCase) Create(role *entity.Role) (*entity.Role, error) {
	// Set timestamps
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	// Create role
	err := u.RoleRepository.CreateRole(role)
	if err != nil {
		return &entity.Role{}, err
	}
	return role, nil
}

// GetByID retrieves a role by ID
func (u *RoleUseCase) GetByID(id uint) (*entity.Role, error) {
	return u.RoleRepository.GetRoleByID(id)
}

// GetByType retrieves a role by type
func (u *RoleUseCase) GetByType(roleType string) ([]*entity.Role, error) {
	return u.RoleRepository.GetRoleByType(roleType)
}

// Update updates a role
func (u *RoleUseCase) Update(role *entity.Role) (*entity.Role, error) {
	existingRole, err := u.RoleRepository.GetRoleByID(role.ID)
	if err != nil {
		return nil, err
	}

	role.CreatedAt = existingRole.CreatedAt
	role.UpdatedAt = time.Now()

	return role,u.RoleRepository.UpdateRole(role)
}

// Delete deletes a role
func (u *RoleUseCase) Delete(id uint) error {
	return u.RoleRepository.DeleteRole(id)
}

// List retrieves all roles
func (u *RoleUseCase) List() ([]*entity.Role, error) {
	return u.RoleRepository.ListRole()
}
