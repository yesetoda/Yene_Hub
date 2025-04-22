package usecases

import (
	"time"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// CountryUseCase defines methods for country business logic
// Now uses schemas for input/output
// PaginationMeta is returned for list endpoints
// Conversion utilities are used internally

type CountryUseCase interface {
	Create(input *schemas.CreateCountryRequest) (*schemas.CountryResponse, error)
	GetByID(id uint) (*schemas.CountryResponse, error)
	GetByName(name string) (*schemas.CountryResponse, error)
	GetByShortCode(shortCode string) (*schemas.CountryResponse, error)
	Update(id uint, input *schemas.UpdateCountryRequest) (*schemas.CountryResponse, error)
	Delete(id uint) error
	List() ([]*schemas.CountryResponse, *schemas.PaginationMeta, error)
}

type countryUseCase struct {
	countryRepo repository.CountryRepository
}

func NewCountryUseCase(countryRepo repository.CountryRepository) CountryUseCase {
	return &countryUseCase{
		countryRepo: countryRepo,
	}
}

func (u *countryUseCase) Create(input *schemas.CreateCountryRequest) (*schemas.CountryResponse, error) {
	country := &entity.Country{
		Name:      input.Name,
		ShortCode: input.Code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	created, err := u.countryRepo.Create(country)
	if err != nil {
		return nil, err
	}
	return entityToCountryResponse(created), nil
}

func (u *countryUseCase) GetByID(id uint) (*schemas.CountryResponse, error) {
	country, err := u.countryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return entityToCountryResponse(country), nil
}

func (u *countryUseCase) GetByName(name string) (*schemas.CountryResponse, error) {
	country, err := u.countryRepo.GetByName(name)
	if err != nil {
		return nil, err
	}
	return entityToCountryResponse(country), nil
}

func (u *countryUseCase) GetByShortCode(shortCode string) (*schemas.CountryResponse, error) {
	country, err := u.countryRepo.GetByShortCode(shortCode)
	if err != nil {
		return nil, err
	}
	return entityToCountryResponse(country), nil
}

func (u *countryUseCase) Update(id uint, input *schemas.UpdateCountryRequest) (*schemas.CountryResponse, error) {
	existing, err := u.countryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Name != nil {
		existing.Name = *input.Name
	}
	if input.Code != nil {
		existing.ShortCode = *input.Code
	}
	existing.UpdatedAt = time.Now()
	updated, err := u.countryRepo.Update(existing)
	if err != nil {
		return nil, err
	}
	return entityToCountryResponse(updated), nil
}

func (u *countryUseCase) Delete(id uint) error {
	return u.countryRepo.Delete(id)
}

func (u *countryUseCase) List() ([]*schemas.CountryResponse, *schemas.PaginationMeta, error) {
	countries, err := u.countryRepo.List()
	if err != nil {
		return nil, nil, err
	}
	resp := make([]*schemas.CountryResponse, 0, len(countries))
	for _, c := range countries {
		resp = append(resp, entityToCountryResponse(c))
	}
	meta := &schemas.PaginationMeta{
		Total:      len(resp),
		Page:       1,
		PageSize:   len(resp),
		TotalPages: 1,
	}
	return resp, meta, nil
}

// entityToCountryResponse converts entity.Country to schemas.CountryResponse
func entityToCountryResponse(c *entity.Country) *schemas.CountryResponse {
	if c == nil {
		return nil
	}
	return &schemas.CountryResponse{
		ID:        c.ID,
		Name:      c.Name,
		Code:      c.ShortCode,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
