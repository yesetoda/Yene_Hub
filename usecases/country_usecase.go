package usecases
import (
	"time"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// CountryUseCase defines methods for country business logic
type CountryUseCase interface {
	Create(country *entity.Country) (*entity.Country, error)
	GetByID(id uint) (*entity.Country, error)
	GetByName(name string) (*entity.Country, error)
	GetByShortCode(shortCode string) (*entity.Country, error)
	Update(country *entity.Country) (*entity.Country, error)
	Delete(id uint) error
	List() ([]*entity.Country, error)
}

// countryUseCase implements CountryUseCase
type countryUseCase struct {
	countryRepo repository.CountryRepository
}

// NewCountryUseCase creates a new CountryUseCase instance
func NewCountryUseCase(countryRepo repository.CountryRepository) CountryUseCase {
	return &countryUseCase{
		countryRepo: countryRepo,
	}
}

// Create creates a new country
func (u *countryUseCase) Create(country *entity.Country) (*entity.Country, error) {
	// Set timestamps
	country.CreatedAt = time.Now()
	country.UpdatedAt = time.Now()

	// Create country
	return u.countryRepo.Create(country)
}

// GetByID retrieves a country by ID
func (u *countryUseCase) GetByID(id uint) (*entity.Country, error) {
	return u.countryRepo.GetByID(id)
}

// GetByName retrieves a country by name
func (u *countryUseCase) GetByName(name string) (*entity.Country, error) {
	return u.countryRepo.GetByName(name)
}

// GetByShortCode retrieves a country by short code
func (u *countryUseCase) GetByShortCode(shortCode string) (*entity.Country, error) {
	return u.countryRepo.GetByShortCode(shortCode)
}

// Update updates a country
func (u *countryUseCase) Update(country *entity.Country) (*entity.Country, error) {
	existingCountry, err := u.countryRepo.GetByID(country.ID)
	if err != nil {
		return nil, err
	}

	country.CreatedAt = existingCountry.CreatedAt
	country.UpdatedAt = time.Now()

	return u.countryRepo.Update(country)
}

// Delete deletes a country
func (u *countryUseCase) Delete(id uint) error {
	return u.countryRepo.Delete(id)
}

// List retrieves all countries
func (u *countryUseCase) List() ([]*entity.Country, error) {
	return u.countryRepo.List()
}
