package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"

	"gorm.io/gorm"
)

// CountryRepository implements the repository.CountryRepository interface
type CountryRepository struct {
	db *gorm.DB
}

// NewCountryRepository creates a new country repository instance
func NewCountryRepository(db *gorm.DB) repository.CountryRepository {
	return &CountryRepository{db: db}
}

// Create creates a new country
func (repo *CountryRepository) Create(country *entity.Country) (*entity.Country, error) {
	result := repo.db.Create(country)
	if result.Error != nil {
		return nil, result.Error
	}
	return country, nil
}

// GetByID retrieves a country by ID
func (repo *CountryRepository) GetByID(id uint) (*entity.Country, error) {
	var country entity.Country
	result := repo.db.First(&country, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}

// GetByName retrieves a country by name
func (repo *CountryRepository) GetByName(name string) (*entity.Country, error) {
	var country entity.Country
	result := repo.db.Where("name = ?", name).First(&country)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}

// GetByShortCode retrieves a country by short code
func (repo *CountryRepository) GetByShortCode(shortCode string) (*entity.Country, error) {
	var country entity.Country
	result := repo.db.Where("short_code = ?", shortCode).First(&country)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}

// Update updates a country
func (repo *CountryRepository) Update(country *entity.Country) (*entity.Country, error) {
	result := repo.db.Save(country)
	if result.Error != nil {
		return nil, result.Error
	}
	return country, nil
}

// Delete deletes a country
func (repo *CountryRepository) Delete(id uint) error {
	result := repo.db.Delete(&entity.Country{}, id)
	return result.Error
}

// List retrieves all countries
func (repo *CountryRepository) List() ([]*entity.Country, error) {
	var countries []*entity.Country
	result := repo.db.Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return countries, nil
}
