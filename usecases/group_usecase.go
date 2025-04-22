package usecases

import (
	"time"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// GroupUseCase defines methods for group business logic
// Now uses schemas for input/output
// PaginationMeta is returned for list endpoints
// Conversion utilities are used internally
type GroupUseCase interface {
	Create(input *schemas.CreateGroupRequest) (*schemas.GroupResponse, error)
	GetByID(id uint) (*schemas.GroupResponse, error)
	GetByName(name string) (*schemas.GroupResponse, error)
	Update(id uint, input *schemas.UpdateGroupRequest) (*schemas.GroupResponse, error)
	Delete(id uint) error
	List() ([]*schemas.GroupResponse, *schemas.PaginationMeta, error)
	GetByCountryID(countryID uint) ([]*schemas.GroupResponse, *schemas.PaginationMeta, error)
}

type groupUseCase struct {
	groupRepo repository.GroupRepository
}

func NewGroupUseCase(groupRepo repository.GroupRepository) GroupUseCase {
	return &groupUseCase{
		groupRepo: groupRepo,
	}
}

func (u *groupUseCase) Create(input *schemas.CreateGroupRequest) (*schemas.GroupResponse, error) {
	group := &entity.Group{
		Name:        input.Name,
		ShortName:   derefString(input.ShortName),
		Description: derefString(input.Description),
		HOAID:       input.HOAID,
		CountryID:   &input.CountryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	created, err := u.groupRepo.Create(group)
	if err != nil {
		return nil, err
	}
	return entityToGroupResponse(created), nil
}

func (u *groupUseCase) GetByID(id uint) (*schemas.GroupResponse, error) {
	group, err := u.groupRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return entityToGroupResponse(group), nil
}

func (u *groupUseCase) GetByName(name string) (*schemas.GroupResponse, error) {
	group, err := u.groupRepo.GetByName(name)
	if err != nil {
		return nil, err
	}
	return entityToGroupResponse(group), nil
}

func (u *groupUseCase) Update(id uint, input *schemas.UpdateGroupRequest) (*schemas.GroupResponse, error) {
	existing, err := u.groupRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Name != nil {
		existing.Name = *input.Name
	}
	if input.ShortName != nil {
		existing.ShortName = derefString(input.ShortName)
	}
	if input.Description != nil {
		existing.Description = derefString(input.Description)
	}
	if input.HOAID != nil {
		existing.HOAID = input.HOAID
	}
	if input.CountryID != nil {
		existing.CountryID = input.CountryID
	}
	existing.UpdatedAt = time.Now()
	updated, err := u.groupRepo.Update(existing)
	if err != nil {
		return nil, err
	}
	return entityToGroupResponse(updated), nil
}

func (u *groupUseCase) Delete(id uint) error {
	return u.groupRepo.Delete(id)
}

func (u *groupUseCase) List() ([]*schemas.GroupResponse, *schemas.PaginationMeta, error) {
	groups, err := u.groupRepo.List()
	if err != nil {
		return nil, nil, err
	}
	resp := make([]*schemas.GroupResponse, 0, len(groups))
	for _, g := range groups {
		resp = append(resp, entityToGroupResponse(g))
	}
	meta := &schemas.PaginationMeta{
		Total:      len(resp),
		Page:       1,
		PageSize:   len(resp),
		TotalPages: 1,
	}
	return resp, meta, nil
}

func (u *groupUseCase) GetByCountryID(countryID uint) ([]*schemas.GroupResponse, *schemas.PaginationMeta, error) {
	groups, err := u.groupRepo.GetGroupsByCountryID(countryID)
	if err != nil {
		return nil, nil, err
	}
	resp := make([]*schemas.GroupResponse, 0, len(groups))
	for _, g := range groups {
		resp = append(resp, entityToGroupResponse(g))
	}
	meta := &schemas.PaginationMeta{
		Total:      len(resp),
		Page:       1,
		PageSize:   len(resp),
		TotalPages: 1,
	}
	return resp, meta, nil
}

func entityToGroupResponse(g *entity.Group) *schemas.GroupResponse {
	if g == nil {
		return nil
	}
	return &schemas.GroupResponse{
		ID:          g.ID,
		Name:        g.Name,
		ShortName:   &g.ShortName,
		Description: &g.Description,
		HOAID:       g.HOAID,
		CountryID:   g.CountryID,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
	}
}

// Helper functions for pointer dereferencing and nil safety
func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefUintPtr(u *uint) uint {
	if u == nil {
		return 0
	}
	return *u
}
