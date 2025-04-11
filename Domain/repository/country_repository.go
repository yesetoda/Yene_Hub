package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// CountryRepository defines methods for country data operations
type CountryRepository interface {
	Create(country *entity.Country) (*entity.Country, error)
	GetByID(id uint) (*entity.Country, error)
	GetByName(name string) (*entity.Country, error)
	GetByShortCode(shortCode string) (*entity.Country, error)
	Update(country *entity.Country) (*entity.Country, error)
	Delete(id uint) error
	List() ([]*entity.Country, error)
}
