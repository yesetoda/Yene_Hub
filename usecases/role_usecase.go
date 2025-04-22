package usecases

import (
	"time"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// RoleUseCase defines methods for role business logic
// Now uses schemas for input/output
// PaginationMeta is returned for list endpoints
// Conversion utilities are used internally
type RoleUseCaseInterface interface {
	Create(input *schemas.CreateRoleRequest) (*schemas.RoleResponse, error)
	GetByID(id uint) (*schemas.RoleResponse, error)
	GetByType(roleType string) ([]*schemas.RoleResponse, error)
	Update(id uint, input *schemas.UpdateRoleRequest) (*schemas.RoleResponse, error)
	Delete(id uint) error
	List() ([]*schemas.RoleResponse, *schemas.PaginationMeta, error)
}

type RoleUseCase struct {
	roleRepo repository.RoleRepository
}

func NewRoleUseCase(roleRepo repository.RoleRepository) *RoleUseCase {
	return &RoleUseCase{
		roleRepo: roleRepo,
	}
}

func (u *RoleUseCase) Create(input *schemas.CreateRoleRequest) (*schemas.RoleResponse, error) {
	role := &entity.Role{
		Type:      input.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := u.roleRepo.CreateRole(role)
	if err != nil {
		return nil, err
	}
	return entityToRoleResponse(role), nil
}

func (u *RoleUseCase) GetByID(id uint) (*schemas.RoleResponse, error) {
	role, err := u.roleRepo.GetRoleByID(id)
	if err != nil {
		return nil, err
	}
	return entityToRoleResponse(role), nil
}

func (u *RoleUseCase) GetByType(roleType string) ([]*schemas.RoleResponse, error) {
	roles, err := u.roleRepo.GetRoleByType(roleType)
	if err != nil {
		return nil, err
	}
	resp := make([]*schemas.RoleResponse, 0, len(roles))
	for _, r := range roles {
		resp = append(resp, entityToRoleResponse(r))
	}
	return resp, nil
}

func (u *RoleUseCase) Update(id uint, input *schemas.UpdateRoleRequest) (*schemas.RoleResponse, error) {
	existing, err := u.roleRepo.GetRoleByID(id)
	if err != nil {
		return nil, err
	}
	if input.Type != nil {
		existing.Type = *input.Type
	}
	existing.UpdatedAt = time.Now()
	err = u.roleRepo.UpdateRole(existing)
	if err != nil {
		return nil, err
	}
	return entityToRoleResponse(existing), nil
}

func (u *RoleUseCase) Delete(id uint) error {
	return u.roleRepo.DeleteRole(id)
}

func (u *RoleUseCase) List() ([]*schemas.RoleResponse, *schemas.PaginationMeta, error) {
	roles, err := u.roleRepo.ListRole()
	if err != nil {
		return nil, nil, err
	}
	resp := make([]*schemas.RoleResponse, 0, len(roles))
	for _, r := range roles {
		resp = append(resp, entityToRoleResponse(r))
	}
	meta := &schemas.PaginationMeta{
		Total:      len(resp),
		Page:       1,
		PageSize:   len(resp),
		TotalPages: 1,
	}
	return resp, meta, nil
}

func entityToRoleResponse(r *entity.Role) *schemas.RoleResponse {
	if r == nil {
		return nil
	}
	return &schemas.RoleResponse{
		ID:        r.ID,
		Type:      r.Type,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
